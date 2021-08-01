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
	//"time"
)

func ChatroomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//アクセスあった時、ルームIDが一致するすべての書き込みをスライスで取得し、テンプレに渡す

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}
		//適当にルームIDを変えると、他の人のルームが覗けるので、メンバのルームしかアクセスできないよう処理
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//リクエスト元のURL文字列からルームIDを取得、数値に変換する
		roomUrl := r.URL.Path
		_roomId := strings.TrimPrefix(roomUrl, "/mypage/chatroom")
		roomId, _ := strconv.Atoi(_roomId)

		//チャットルームDBに接続する
		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		//ルームIDに一致するチャットルームを取得する
		selectedChatroom := query.SelectChatroomById(roomId, dbChtrm)

		//アクセスしてきたユーザーが、チャットルームのメンバーでなかった場合、アクセスを拒否する
		if selectedChatroom.UserId != userSessionVar && selectedChatroom.Member != userSessionVar {
			fmt.Fprintf(w, "ルームにアクセスする権限がありません")
			return
		}

		t := template.Must(template.ParseFiles("./templates/mypage/chatroom.html"))
		t.ExecuteTemplate(w, "chatroom.html", nil)

	case "POST":
		//チャットルーム削除ボタンが押された時の処理

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		//ルーム削除ボタンが押されPOSTがあった時、当該ルームを削除する
		if r.FormValue("delete-room") == "このルームを削除する" {

			//リクエスト元のURL文字列からルームIDを取得、数値に変換する
			roomUrl := r.URL.Path
			_roomId := strings.TrimPrefix(roomUrl, "/mypage/chatroom")
			roomId, _ := strconv.Atoi(_roomId)

			//チャットルームDBに接続する
			dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbChtrm.Close()

			//ルームIDと一致するチャットルームを削除する
			query.DeleteChatroomById(roomId, dbChtrm)

			t := template.Must(template.ParseFiles("./templates/mypage/chatroomdeleted.html"))
			t.ExecuteTemplate(w, "chatroomdeleted.html", nil)
		}
	}
}
