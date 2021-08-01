//マイページへアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"time"
	"web-server/query"
	"web-server/sessions"
)

func MypageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		/*アクセスがあった時の処理。自身で作成したルームと他人が作成したルームを別々に取得して、
		ルーム一覧のスライスをつくってテンプレに渡す*/

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		//チャットルームDBに接続する
		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//アクセスしたユーザーが作成したチャットルームを全て取得し、スライスで返す
		chatroomsFromUserId := query.SelectAllChatroomsByUserId(userSessionVar, dbChtrm)

		//アクセスしたユーザーがメンバーとして参加しているチャットルームを全て取得し、スライスで返す
		chatroomsFromMember := query.SelectAllChatroomsByMember(userSessionVar, dbChtrm)

		//以上のスライスをまとめ、マイページでリンク表示するためデータを用意、ExecuteTemplateに渡す
		var Links = append(chatroomsFromUserId, chatroomsFromMember...)
		t := template.Must(template.ParseFiles("./templates/mypage.html"))
		t.ExecuteTemplate(w, "mypage.html", Links)

	case "POST":
		/*新しいルーム作成のポストがあった時の処理。ルーム名と相手メンバーを指定する。
		同名のルーム名は、相手メンバー異なる場合のみ有効。*/

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/sessionexpired.html"))
			t.ExecuteTemplate(w, "sessionexpired.html", nil)
			return
		}

		//新規チャットルームのため、チャットルーム構造体を初期化する
		newchatroom := new(query.Chatroom)

		//フォームに入力された値を取得する
		roomname := r.FormValue("roomName")
		newchatroom.Member = r.FormValue("memberName")

		//メンバーまたはルーム名が入力されていない場合は処理を中断する
		if roomname == "" || newchatroom.Member == "" {
			return
		}

		//ルーム名に含まれるメタ文字をエスケープする
		newchatroom.RoomName = regexp.QuoteMeta(roomname)

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//自分自身をメンバーにしていた場合は処理を中断する
		if newchatroom.Member == userSessionVar {
			return
		}

		//ユーザーDBに接続する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//全てのユーザーを取得し、相手ユーザーが登録されていない場合は処理を中断する
		users := query.SelectAllUser(dbUsr)
		userIdExist := query.ContainsUserName(users, newchatroom.Member)
		if !userIdExist {
			return
		}

		//チャットルームDBに接続する
		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		//新たなチャットルームをDBに登録する
		newChtrmInsertedToDb := query.InsertChatroom(userSessionVar, newchatroom.RoomName, newchatroom.Member, dbChtrm)
		if newChtrmInsertedToDb {

			//チャットルームDBに再度接続する
			dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbChtrm.Close()

			//以下、WebSocketのハンドシェイクのための初期投稿処理

			//新規チャットのため、チャット構造体を初期化する
			newChat := new(query.Chat)
			newChat.Chat = "NEW ROOM CREATED"

			//ルームIDを取得する必要があるため、DBから改めて登録したルーム情報を取得する
			createdChatroom := query.SelectChatroomByUserAndRoomNameAndMember(userSessionVar, newchatroom.RoomName, newchatroom.Member, dbChtrm)

			//チャットインスタンスに、ルームID、ユーザーID、ルーム名、メンバーを代入する
			newChat.Chatroom.Id = createdChatroom.Id
			newChat.Chatroom.UserId = userSessionVar
			newChat.Chatroom.RoomName = createdChatroom.RoomName
			newChat.Chatroom.Member = createdChatroom.Member

			//チャット投稿時間を取得する
			newChat.PostDt = time.Now().UTC().Round(time.Second)

			//チャットをDBに登録
			posted := query.InsertChat(newChat.Chatroom.Id, newChat.Chatroom.UserId, newChat.Chatroom.RoomName, newChat.Chatroom.Member, newChat.Chat, newChat.PostDt, dbChtrm)
			if posted == true {
				fmt.Println("初期投稿成功")
				return
			} else {
				fmt.Println("初期投稿失敗")
				return
			}
		} else {
			fmt.Println("ルーム作成失敗")
		}
	}
}
