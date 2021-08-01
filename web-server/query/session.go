//セッション情報を扱うクエリパッケージ
package query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"web-server/config"
)

var ConStrSession string

func init() {
	confDbSession, err := config.ReadConfDbSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	ConStrSession = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDbSession.User, confDbSession.Pass, confDbSession.Host, confDbSession.Port, confDbSession.DbName, confDbSession.Charset)
}

// セッション情報をDBに登録する関数
func InsertSession(sessionId string, sessionValue string, db *sql.DB) bool {

	stmt, err := db.Prepare("INSERT INTO SESSIONS(SESSION_ID,SESSION_VALUE) VALUES(?,?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionId, sessionValue)
	if err != nil {
		return false
	} else {
		return true
	}
}

//ユーザーIDと一致するユーザー情報をDBから取得する関数
func SelectSessionBySessionId(sessionId string, db *sql.DB) (sessionValue string) {

	err := db.QueryRow("SELECT SESSION_VALUE FROM SESSION_ID WHERE SESSION_ID = ?", sessionId).Scan(&sessionValue)
	if err != nil {
		return
	}
	return
}

//ユーザーIDに一致するユーザーをDBから削除する関数
func DeleteSessionBySessionId(sessionId string, db *sql.DB) bool {

	stmt, err := db.Prepare("DELETE FROM SESSIONS WHERE SESSION_ID = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionId)
	if err != nil {
		return false
	} else {
		return true
	}
}
