//管理ページからログアウトした時のハンドラ
package routers

import (
	"html/template"
	"net/http"
	"web-server/sessions"
)

func AdminLogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		session.Manager.DeleteSessionFromStore(w, r)

		t := template.Must(template.ParseFiles("./templates/admin/adminlogout.html"))
		t.ExecuteTemplate(w, "adminlogout.html", nil)
	}
}
