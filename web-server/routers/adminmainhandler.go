//管理ページ用ハンドラ(メインページ)
package routers

import (
	//"golang.org/x/crypto/bcrypt"
	//"fmt"
	//"goserver/sessions"
	"html/template"
	"net/http"
)

func AdminMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// if ok := session.Manager.SessionIdCheck(w, r); !ok {
		// 	t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
		// 	t.ExecuteTemplate(w, "sessionexpired.html", nil)
		// 	return
		// }

		t := template.Must(template.ParseFiles("./templates/admin/main.html"))
		t.ExecuteTemplate(w, "main.html", nil)
		return
	}
}
