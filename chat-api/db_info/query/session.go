//セッション情報を扱うクエリパッケージ
package query

import (
	"chat-api/db_info/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var ConStrSession string

func init() {
	confDbSession, err := config.ReadConfDbSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	ConStrSession = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDbSession.User, confDbSession.Pass, confDbSession.Host, confDbSession.Port, confDbSession.DbName, confDbSession.Charset)
}

//ユーザーIDと一致するユーザー情報をdbから取得する関数
func SelectSessionBySessionId(sessionId string, db *sql.DB) string {

	var sessionValue string

	err := db.QueryRow("SELECT SESSION_VALUE FROM SESSIONS WHERE SESSION_ID = ?", sessionId).Scan(&sessionValue)
	if err != nil {
		fmt.Println(err.Error())
	}
	return sessionValue
}
