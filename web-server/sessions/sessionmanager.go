package session

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"net/url"
	"time"
	"web-server/query"
)

//セッション構造体
type Session struct {
	sid          string
	timeAccessed time.Time
	SessionValue map[string]string
	Manager      *MANAGER
}

//セッションマネージャ構造体
type MANAGER struct {
	SessionStore map[interface{}]*Session
	CookieName   string
	maxlifetime  int64
}

var Manager *MANAGER

//サーバー起動時にセッションマネージャも初期化
func init() {
	Manager = NewManager("cookieName", 3600)

}

//init()でセッションマネージャ初期化時に使う関数
func NewManager(cookieName string, maxlifetime int64) *MANAGER {
	database := make(map[interface{}]*Session)
	return &MANAGER{SessionStore: database, CookieName: cookieName, maxlifetime: maxlifetime}
}

//新規セッションを開始する関数(セッション変数はマネージャのDatabaseマップに保存)
func (Manager *MANAGER) SessionStart(w http.ResponseWriter, r *http.Request, userId string) (session *Session) {
	cookie, err := r.Cookie(Manager.CookieName)
	if err != nil || cookie.Value == "" {

		//セッションID発行
		sid := Manager.NewSessionId()

		//発行したセッションIDを元にセッション生成
		session := Manager.NewSession(sid, userId)

		//cookieを取得するため、HttpOnlyをfalseに設定中。https環境では、Secure属性
		cookie := http.Cookie{Name: Manager.CookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: false, MaxAge: int(Manager.maxlifetime)}

		//セッションDBに接続する
		dbSession, err := sql.Open("mysql", query.ConStrSession)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbSession.Close()

		//セッションDBにセッションID・ユーザーIDをのペアを保存する
		query.InsertSession(sid, userId, dbSession)

		//クライアントのブラウザにcookieをセットする
		http.SetCookie(w, &cookie)

		//セッションマネージャのストアにセッションIDをキーにセッションを保持
		Manager.SessionStore[sid] = session
	}
	return
}

//新規セッションIDを発行する関数(SessionStart関数で使用)
func (Manager *MANAGER) NewSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//新規セッション(マップ)を作成するための関数(SessionStart関数で使用)
func (Manager *MANAGER) NewSession(sid string, userId string) (session *Session) {
	sv := make(map[string]string)

	sv["userId"] = userId
	return &Session{sid: sid, timeAccessed: time.Now(), SessionValue: sv}
}

//クライアントのクッキーid(セッションid)が、マネージャのDatabaseマップに登録されてるかチェックする関数
func (Manager *MANAGER) SessionIdCheck(w http.ResponseWriter, r *http.Request) bool {
	//クライアントのcookieを取得する
	clientCookie, err := r.Cookie(Manager.CookieName)
	if err != nil {
		return false
	} else {
		//クライアントのcookieからセッションIDを取得する
		clientSid, _ := url.QueryUnescape(clientCookie.Value)
		if _, ok := Manager.SessionStore[clientSid]; ok {
			//クライアントのセッションの生成時間＋セッション寿命と現在の時刻を比較する
			if (Manager.SessionStore[clientSid].timeAccessed.Unix() + Manager.maxlifetime) > time.Now().Unix() {
				//セッションの生成時間＋セッション寿命＞現在の時刻のときは、生成時間を現在の時刻に更新
				Manager.SessionStore[clientSid].timeAccessed = time.Now()
				return true
			} else {
				//セッションの生成時間＋セッション寿命＜現在の時刻のときはfalseを返す
				return false
			}
		} else {
			return false
		}
	}

}

//ログアウト時、セッションを削除マネージャのDatabaseマップから削除する関数
func (Manager *MANAGER) DeleteSessionFromStore(w http.ResponseWriter, r *http.Request) error {
	//クライアントのcookieを取得する
	clientCookie, err := r.Cookie(Manager.CookieName)
	if err != nil {
		fmt.Println(err.Error())
	}
	//クライアントのcookieからセッションIDを取得する
	clientSid, _ := url.QueryUnescape(clientCookie.Value)

	//セッションマネージャのストアからセッションIDに一致するセッションを削除する
	delete(Manager.SessionStore, clientSid)

	//セッションDBに接続する
	dbSession, err := sql.Open("mysql", query.ConStrSession)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbSession.Close()

	//セッションDBからセッションIDに一致するセッションを削除する
	query.DeleteSessionBySessionId(clientSid, dbSession)

	//クライアントのcookieにMaxAge=-1を設定し、クライアントcookieを削除
	clientCookie.MaxAge = -1
	http.SetCookie(w, clientCookie)
	return nil
}
