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

//メモリ上でセッションを管理するための、セッションマネージャ構造体
type MANAGER struct {
	SessionStore map[interface{}]*Session
	CookieName   string
	maxlifetime  int64
}

var Manager *MANAGER

func init() {
	Manager = NewManager("cookieName", 3600)

}

func NewManager(cookieName string, maxlifetime int64) *MANAGER {
	database := make(map[interface{}]*Session)
	return &MANAGER{SessionStore: database, CookieName: cookieName, maxlifetime: maxlifetime}
}

func (Manager *MANAGER) SessionStart(w http.ResponseWriter, r *http.Request, userId string) (session *Session) {
	cookie, err := r.Cookie(Manager.CookieName)
	if err != nil || cookie.Value == "" {

		sid := Manager.GenerateNewSessionId()

		session := Manager.CreateNewSession(sid, userId)

		//管理ページからのアクセスの際は、Path属性を/adminに設定する必要があるため、ユーザー名で分岐処理する
		if userId == "admin" {
			//cookieを取得するため、HttpOnlyをfalseに設定中。https環境では、Secure属性を設定する
			cookie := http.Cookie{Name: Manager.CookieName, Value: url.QueryEscape(sid), Path: "/admin", HttpOnly: false, MaxAge: int(Manager.maxlifetime)}
			http.SetCookie(w, &cookie)
		} else {
			cookie := http.Cookie{Name: Manager.CookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: false, MaxAge: int(Manager.maxlifetime)}
			http.SetCookie(w, &cookie)
		}

		dbSession, err := sql.Open("mysql", query.ConStrSession)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbSession.Close()

		query.InsertSession(sid, userId, dbSession)

		Manager.SessionStore[sid] = session
		fmt.Println(Manager.SessionStore[sid])
	}
	return
}

func (Manager *MANAGER) GenerateNewSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (Manager *MANAGER) CreateNewSession(sid string, userId string) (session *Session) {
	sv := make(map[string]string)

	sv["userId"] = userId
	return &Session{sid: sid, timeAccessed: time.Now(), SessionValue: sv}
}

//クライアントのクッキーid(セッションid)が、マネージャのDatabaseマップに登録されてるかチェックする関数
func (Manager *MANAGER) SessionIdCheck(w http.ResponseWriter, r *http.Request) bool {
	clientCookie, err := r.Cookie(Manager.CookieName)
	if err != nil {
		return false
	} else {
		clientSid, _ := url.QueryUnescape(clientCookie.Value)

		if _, idExist := Manager.SessionStore[clientSid]; idExist {
			if (Manager.SessionStore[clientSid].timeAccessed.Unix() + Manager.maxlifetime) > time.Now().Unix() {
				Manager.SessionStore[clientSid].timeAccessed = time.Now()
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}
}

func (Manager *MANAGER) DeleteSessionFromStore(w http.ResponseWriter, r *http.Request) error {

	clientCookie, err := r.Cookie(Manager.CookieName)
	if err != nil {
		return nil
	}

	clientSid, _ := url.QueryUnescape(clientCookie.Value)

	delete(Manager.SessionStore, clientSid)

	dbSession, err := sql.Open("mysql", query.ConStrSession)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbSession.Close()

	query.DeleteSessionBySessionId(clientSid, dbSession)

	clientCookie.MaxAge = -1
	http.SetCookie(w, clientCookie)
	return nil
}
