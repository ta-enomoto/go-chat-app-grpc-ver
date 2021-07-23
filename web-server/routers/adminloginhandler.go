//管理ページ用ハンドラ(ログイン)
package routers

import (
	//"golang.org/x/crypto/bcrypt"
	"fmt"
	"goserver/sessions"
	"html/template"
	"net/http"
)

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.Must(template.ParseFiles("./templates/admin/login.html"))
		t.ExecuteTemplate(w, "login.html", nil)
		return

	//管理者ID・パスワードがポストされた時
	case "POST":
		adminUserId := "admin"
		adminPassword := "pass"

		userid := r.FormValue("userId")
		password := r.FormValue("password")

		if userid == "" || password == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません")
			return
		}

		if userid != adminUserId {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		if password != adminPassword {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		session.Manager.SessionStart(w, r, adminUserId)
		http.Redirect(w, r, "/admin/main", 301)
	}
}
