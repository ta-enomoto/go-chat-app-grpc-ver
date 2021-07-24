//管理ページ用ハンドラ(チャットルーム一覧ページ)
package routers

import (
	//"golang.org/x/crypto/bcrypt"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goserver/query"
	//"goserver/sessions"
	"html/template"
	"net/http"
)

func AdminChatroomsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// if ok := session.Manager.SessionIdCheck(w, r); !ok {
		// 	t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
		// 	t.ExecuteTemplate(w, "sessionexpired.html", nil)
		// 	return
		// }

		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		AllChatrooms := query.SelectAllChatrooms(dbChtrm)

		t := template.Must(template.ParseFiles("./templates/admin/chatrooms.html"))
		t.ExecuteTemplate(w, "chatrooms.html", AllChatrooms)
		return
	}
}
