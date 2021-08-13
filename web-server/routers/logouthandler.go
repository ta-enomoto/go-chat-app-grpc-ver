//マイページからログアウトした時のハンドラ
package routers

import (
	"html/template"
	"net/http"
	"web-server/sessions"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		session.Manager.DeleteSessionFromStore(w, r)

		t := template.Must(template.ParseFiles("./templates/logout.html"))
		t.ExecuteTemplate(w, "logout.html", nil)
	}
}
