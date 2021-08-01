//あるチャットルームの個別管理ページにアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"web-server/query"
	"web-server/sessions"
)

func AdminChatroomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//ユーザーIDがadmin以外のユーザーのアクセスを拒否する
		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		t := template.Must(template.ParseFiles("./templates/admin/adminchatroom/adminchatroom.html"))
		t.ExecuteTemplate(w, "adminchatroom.html", nil)

	case "POST":
		//チャットルーム削除ボタンが押された時の処理

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//ユーザーIDがadmin以外のユーザーのアクセスを拒否する
		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		//ルーム削除ボタンが押されPOSTがあった時、当該ルームを削除する
		if r.FormValue("delete-room") == "このルームを削除する" {

			//リクエスト元のURL文字列からルームIDを取得、数値に変換する
			roomUrl := r.URL.Path
			_roomId := strings.TrimPrefix(roomUrl, "/admin/chatrooms/chatroom")
			roomId, _ := strconv.Atoi(_roomId)

			//チャットルームDBに接続する
			dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbChtrm.Close()

			//ルームIDに一致するチャットルームを削除する
			query.DeleteChatroomById(roomId, dbChtrm)

			t := template.Must(template.ParseFiles("./templates/admin/adminchatroom/adminchatroomdeleted.html"))
			t.ExecuteTemplate(w, "adminchatroomdeleted.html", nil)
		}
	}
}
