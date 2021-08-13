//管理者ログインページにアクセスがあった時のハンドラ
package routers

import (
	"fmt"
	"html/template"
	"net/http"
	"web-server/sessions"
)

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.Must(template.ParseFiles("./templates/admin/adminlogin.html"))
		t.ExecuteTemplate(w, "adminlogin.html", nil)

	case "POST":
		//ログインボタンが押され管理者ID・パスワードがPOSTされた時の処理

		//DBのボリュームに初期値として入れる方法は、ハッシュ化の関係で出来ないため、管理者ID・パスワードは以下で固定する。
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

		session.Manager.SessionStart(w, r, userid)
		http.Redirect(w, r, "/admin/main", 301)
	}
}
