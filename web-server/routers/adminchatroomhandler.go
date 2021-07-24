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
	"strconv"
	"strings"
)

func AdminChatroomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		t := template.Must(template.ParseFiles("./templates/admin/chatrooms/chatroom.html"))
		t.ExecuteTemplate(w, "chatroom.html", nil)
		return

	case "POST":
		// if ok := session.Manager.SessionIdCheck(w, r); !ok {
		// 	t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
		// 	t.ExecuteTemplate(w, "sessionexpired.html", nil)
		// 	return
		// }
		if r.FormValue("delete-room") == "このルームを削除する" {
			roomUrl := r.URL.Path
			_roomId := strings.TrimPrefix(roomUrl, "/admin/chatrooms/chatroom")
			roomId, _ := strconv.Atoi(_roomId)

			dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbChtrm.Close()

			query.DeleteChatroomById(roomId, dbChtrm)

			t := template.Must(template.ParseFiles("./templates/mypage/chatroomdeleted.html"))
			t.ExecuteTemplate(w, "chatroomdeleted.html", nil)
		}
	}
}
