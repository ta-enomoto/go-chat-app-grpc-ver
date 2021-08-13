//ユーザー登録ページにアクセスがあったときのハンドラ
package routers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"regexp"
	"web-server/query"
)

func ResistrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.Must(template.ParseFiles("./templates/resistration.html"))
		t.ExecuteTemplate(w, "resistration.html", nil)

	case "POST":
		//新規登録ユーザーがPOSTされた時の処理

		newUser := new(query.User)

		newUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")

		if newUser.UserId == "" || psw_string == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません。")
			return
		}

		escapeStrings := regexp.MustCompile(`\?|\$|\&|\=|\-|\>|\<|\+|\;|\:|\*|\||\'`)
		if escapeStrings.MatchString(newUser.UserId) {
			fmt.Fprintf(w, "使用できない文字が含まれています。")
			return
		}
		if escapeStrings.MatchString(psw_string) {
			fmt.Fprintf(w, "使用できない文字が含まれています。")
			return
		}

		//ユーザー名adminは管理者IDとして使用不可
		if newUser.UserId == "admin" {
			fmt.Fprintf(w, "使用できないユーザー名です。")
			return
		}

		newUser.Password = []byte(psw_string)

		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		users := query.SelectAllUser(dbUsr)
		userIdAlreadyExists := query.ContainsUserName(users, newUser.UserId)
		if userIdAlreadyExists {
			fmt.Fprintf(w, "既に登録されているIDです。")
			return
		}

		//ユーザーをDBに登録する(パスワードのハッシュ化はChangePassword関数内で行う)
		insertedUser := query.InsertUser(newUser.UserId, newUser.Password, dbUsr)
		if insertedUser {
			t := template.Must(template.ParseFiles("./templates/resistrationcompleted.html"))
			t.ExecuteTemplate(w, "resistrationcompleted.html", nil)
		}
	}
}
