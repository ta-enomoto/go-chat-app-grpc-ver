//ユーザー一覧管理ページにアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"net/url"
	"web-server/query"
	"web-server/sessions"
)

func AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
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

		//ユーザーDBに接続する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//全てのユーザーを取得する
		AllUsers := query.SelectAllUser(dbUsr)

		//全ユーザーをブラウザで表示するため、ExecuteTemplateの引数に指定する
		t := template.Must(template.ParseFiles("./templates/admin/adminusers.html"))
		t.ExecuteTemplate(w, "adminusers.html", AllUsers)
	}
}
