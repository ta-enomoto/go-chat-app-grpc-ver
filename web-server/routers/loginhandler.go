//ログインページにアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"web-server/query"
	"web-server/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.Must(template.ParseFiles("./templates/login.html"))
		t.ExecuteTemplate(w, "login.html", nil)
		return

	case "POST":
		//ログインID・パスワードがポストされた時の処理

		//アクセスしてきたユーザーのため、ユーザー構造体を初期化する
		accessingUser := new(query.User)

		//フォームに入力された値を取得する
		accessingUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")

		//フォームに何も入力されていない時の処理(ブラウザ側でもチェック有り)
		if accessingUser.UserId == "" || psw_string == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません。")
			return
		}

		//ハッシュ化のため、パスワードをbyte型に変換する
		accessingUser.Password = []byte(psw_string)

		//ユーザーDBに接続する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//ユーザーDBから、アクセスがあったユーザーの情報を取得する
		user := query.SelectUserById(accessingUser.UserId, dbUsr)

		//ユーザーIDが間違っていた場合の処理
		if user.UserId == "" {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
			return
		}

		//パスワードが間違っていた時の処理
		err = bcrypt.CompareHashAndPassword(user.Password, accessingUser.Password)
		if err != nil {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		//セッションを生成
		session.Manager.SessionStart(w, r, accessingUser.UserId)

		//マイページにリダイレクト
		http.Redirect(w, r, "/mypage", 301)
	}
}
