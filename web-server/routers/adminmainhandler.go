//メイン管理ページにアクセスがあった時のハンドラ
package routers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"web-server/sessions"
)

func AdminMainHandler(w http.ResponseWriter, r *http.Request) {
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

		t := template.Must(template.ParseFiles("./templates/admin/adminmain.html"))
		t.ExecuteTemplate(w, "adminmain.html", nil)
	}
}
