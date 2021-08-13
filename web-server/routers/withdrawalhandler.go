//退会ページにアクセスがあった時のハンドラ
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

func WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		t := template.Must(template.ParseFiles("./templates/withdrawal.html"))
		t.ExecuteTemplate(w, "withdrawal.html", nil)

	case "POST":
		//削除するユーザーのID・パスワードがポストされたときの処理
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		deleteUser := new(query.User)
		deleteUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")
		deleteUser.Password = []byte(psw_string)

		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		user := query.SelectUserById(deleteUser.UserId, dbUsr)

		if user.UserId == "" {
			fmt.Fprintf(w, "IDが間違っています。")
			return
		}

		err = bcrypt.CompareHashAndPassword(user.Password, deleteUser.Password)
		if err != nil {
			fmt.Fprintf(w, "パスワードが間違っています。")
			return
		}

		userDeletedFromDb := query.DeleteUserById(deleteUser.UserId, dbUsr)
		if userDeletedFromDb {
			session.Manager.DeleteSessionFromStore(w, r)

			t := template.Must(template.ParseFiles("./templates/withdrawalcompleted.html"))
			t.ExecuteTemplate(w, "withdrawalcompleted.html", nil)
		}
	}
}
