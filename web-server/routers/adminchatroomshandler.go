//管理ページ用ハンドラ(チャットルーム一覧ページ)
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
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		AllChatrooms := query.SelectAllChatrooms(dbChtrm)

		t := template.Must(template.ParseFiles("./templates/admin/adminchatrooms.html"))
		t.ExecuteTemplate(w, "adminchatrooms.html", AllChatrooms)
		return
	}
}
