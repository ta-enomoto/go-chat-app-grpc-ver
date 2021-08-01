//チャットルーム一覧管理ページにアクセスがあった時のハンドラ
package routers

import (
	//"golang.org/x/crypto/bcrypt"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"net/url"
	"web-server/query"
	"web-server/sessions"
)

func AdminChatroomsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//ユーザーIDがadmin以外のユーザーのアクセスを拒否する
		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		//チャットルームDBに接続する
		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		//DB上の全てのチャットルームを取得する
		AllChatrooms := query.SelectAllChatrooms(dbChtrm)

		//全チャットルームをブラウザで表示するため、ExecuteTemplateの引数に指定する
		t := template.Must(template.ParseFiles("./templates/admin/adminchatrooms.html"))
		t.ExecuteTemplate(w, "adminchatrooms.html", AllChatrooms)
	}
}
