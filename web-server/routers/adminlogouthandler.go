package routers

import (
	"goserver/sessions"
	"html/template"
	"net/http"
)

func AdminLogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		session.Manager.DeleteSessionFromStore(w, r)
		t := template.Must(template.ParseFiles("./templates/admin/adminlogout.html"))
		t.ExecuteTemplate(w, "adminlogout.html", nil)
	}
}
