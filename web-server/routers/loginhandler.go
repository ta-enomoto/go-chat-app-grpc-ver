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

		accessingUser := new(query.User)

		accessingUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")

		if accessingUser.UserId == "" || psw_string == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません。")
			return
		}

		accessingUser.Password = []byte(psw_string)

		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		user := query.SelectUserById(accessingUser.UserId, dbUsr)

		if user.UserId == "" {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
			return
		}

		err = bcrypt.CompareHashAndPassword(user.Password, accessingUser.Password)
		if err != nil {
			fmt.Fprintf(w, "IDまたはパスワードが間違っています。")
		}

		session.Manager.SessionStart(w, r, accessingUser.UserId)
		http.Redirect(w, r, "/mypage", 301)
	}
}
