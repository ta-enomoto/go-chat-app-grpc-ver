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

		//管理者ID・パスワードは以下で固定する
		adminUserId := "admin"
		adminPassword := "pass"

		//フォームに入力された値を取得する
		userid := r.FormValue("userId")
		password := r.FormValue("password")

		//フォームに何も入力されていない時の処理(ブラウザ側でもチェック有り)
		if userid == "" || password == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません")
			return
		}

		//ユーザーIDが誤っている場合の処理
		if userid != adminUserId {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		//パスワードが誤っている場合の処理
		if password != adminPassword {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		//セッションを生成
		session.Manager.SessionStart(w, r, userid)

		//管理メインページにリダイレクト
		http.Redirect(w, r, "/admin/main", 301)
	}
}
