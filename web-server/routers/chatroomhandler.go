//個別のチャットルームにアクセスがあったときのハンドラ
package routers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"web-server/query"
	"web-server/sessions"
)

func ChatroomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//アクセスあった時、ルームIDが一致するすべての書き込みをスライスで取得し、テンプレに渡す

		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		//適当にルームIDを変えて他の人のルームが覗けないよう、メンバになっているルームしかアクセスできないよう処理
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		roomUrl := r.URL.Path
		_roomId := strings.TrimPrefix(roomUrl, "/mypage/chatroom")
		roomId, _ := strconv.Atoi(_roomId)

		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		selectedChatroom := query.SelectChatroomById(roomId, dbChtrm)

		if selectedChatroom.UserId != userSessionVar && selectedChatroom.Member != userSessionVar {
			fmt.Fprintf(w, "ルームにアクセスする権限がありません")
			return
		}

		t := template.Must(template.ParseFiles("./templates/mypage/chatroom.html"))
		t.ExecuteTemplate(w, "chatroom.html", nil)

	case "POST":
		//チャットルーム削除ボタンが押された時の処理

		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		if r.FormValue("delete-room") == "このルームを削除する" {

			roomUrl := r.URL.Path
			_roomId := strings.TrimPrefix(roomUrl, "/mypage/chatroom")
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
