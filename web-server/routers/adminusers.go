//管理ページ用ハンドラ(ユーザー一覧ページ)
package routers

import (
	//"golang.org/x/crypto/bcrypt"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goserver/query"
	"goserver/sessions"
	"html/template"
	"net/http"
)

func AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		AllUsers := query.SelectAllUser(dbUsr)

		t := template.Must(template.ParseFiles("./templates/admin/users.html"))
		t.ExecuteTemplate(w, "users.html", AllUsers)
		return
	}
}
