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

		//新規ユーザーのため、ユーザー構造体を初期化する
		newUser := new(query.User)

		//フォームに入力された値を取得する
		newUser.UserId = r.FormValue("userId")
		psw_string := r.FormValue("password")

		//フォームに何も入力されていない時の処理(ブラウザ側でもチェック有り)
		if newUser.UserId == "" || psw_string == "" {
			fmt.Fprintf(w, "IDまたはパスワードが入力されていません。")
			return
		}

		//使用不可の特殊記号のチェックを行う
		escapeStrings := regexp.MustCompile(`\?|\$|\&|\=|\-|\>|\<|\+|\;|\:|\*|\||\'`)
		if escapeStrings.MatchString(newUser.UserId) {
			fmt.Fprintf(w, "使用できない文字が含まれています。")
			return
		}
		if escapeStrings.MatchString(psw_string) {
			fmt.Fprintf(w, "使用できない文字が含まれています。")
			return
		}

		//ユーザー名adminは禁止
		if newUser.UserId == "admin" {
			fmt.Fprintf(w, "使用できないユーザー名です。")
			return
		}

		//パスワードはハッシュ化のため、byte型に変換する
		newUser.Password = []byte(psw_string)

		//ユーザーDBに接続する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//ユーザー名が重複していないかチェックするため、全ユーザーを取得する
		users := query.SelectAllUser(dbUsr)

		//ユーザー名の重複をチェックする
		userIdAlreadyExists := query.ContainsUserName(users, newUser.UserId)
		if userIdAlreadyExists {
			fmt.Fprintf(w, "既に登録されているIDです。")
			return
		}

		//ユーザーをDBに登録する
		insertedUser := query.InsertUser(newUser.UserId, newUser.Password, dbUsr)
		if insertedUser {
			t := template.Must(template.ParseFiles("./templates/resistrationcompleted.html"))
			t.ExecuteTemplate(w, "resistrationcompleted.html", nil)
		}
	}
}
