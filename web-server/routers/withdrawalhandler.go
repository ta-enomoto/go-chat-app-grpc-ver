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
		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		t := template.Must(template.ParseFiles("./templates/withdrawal.html"))
		t.ExecuteTemplate(w, "withdrawal.html", nil)

	//削除するユーザーのID・パスワードがポストされたときの処理
	case "POST":
		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		//新規ユーザーインスタンスを生成する
		deleteUser := new(query.User)
		deleteUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")
		deleteUser.Password = []byte(psw_string)

		//フォームに入力された値を取得する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//ユーザーIDと一致するユーザー情報を取得する
		user := query.SelectUserById(deleteUser.UserId, dbUsr)

		//ユーザーIDが間違っていた場合の処理
		if user.UserId == "" {
			fmt.Fprintf(w, "IDが間違っています。")
			return
		}

		//パスワードが間違っていた場合の処理
		err = bcrypt.CompareHashAndPassword(user.Password, deleteUser.Password)
		if err != nil {
			fmt.Fprintf(w, "パスワードが間違っています。")
			return
		}

		//当該ユーザーをDBから削除する
		userDeletedFromDb := query.DeleteUserById(deleteUser.UserId, dbUsr)
		if userDeletedFromDb {
			//当該セッションを削除する
			session.Manager.DeleteSessionFromStore(w, r)

			t := template.Must(template.ParseFiles("./templates/withdrawalcompleted.html"))
			t.ExecuteTemplate(w, "withdrawalcompleted.html", nil)
		}
	}
}
