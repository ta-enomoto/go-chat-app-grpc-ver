/*サーバーの起動＆アクセスに対するルーティング*/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
	"net/http"
	"regexp"
	"web-server/routers"
	"web-server/wsserver"
)

type MyMux struct{}

/*個別のチャットルームへルーティングするのに正規表現を使用する都合上、他のルーティングもすべて正規表現を使用
(文字列でのswitchと正規表現との一致によるswitchが混在できない)*/

func (mux MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var login = regexp.MustCompile(`^/login$`)
	var mypage = regexp.MustCompile(`^/mypage$`)
	var resistration = regexp.MustCompile(`^/resistration$`)
	var logout = regexp.MustCompile(`^/logout$`)
	var withdrawal = regexp.MustCompile(`^/withdrawal$`)
	var dirOfChatroom = regexp.MustCompile(`/mypage/chatroom[0-9]*`)
	var websocketSever = regexp.MustCompile(`^/wsserver$`)
	var publicFiles = regexp.MustCompile(`/public/.*`)
	var adminLogin = regexp.MustCompile(`^/admin/login$`)
	var adminMain = regexp.MustCompile(`^/admin/main$`)
	var adminUsers = regexp.MustCompile(`^/admin/users$`)
	var adminChatrooms = regexp.MustCompile(`^/admin/chatrooms$`)
	var adminUser = regexp.MustCompile(`^/admin/users/.*`)
	var adminChatroom = regexp.MustCompile(`^/admin/chatroom[0-9]*`)
	var adminLogout = regexp.MustCompile(`^/admin/logout$`)
	url := r.URL.Path

	switch {
	case login.MatchString(url):
		routers.LoginHandler(w, r)

	case mypage.MatchString(url):
		routers.MypageHandler(w, r)

	case resistration.MatchString(url):
		routers.ResistrationHandler(w, r)

	case logout.MatchString(url):
		routers.LogoutHandler(w, r)

	case withdrawal.MatchString(url):
		routers.WithdrawalHandler(w, r)

	case dirOfChatroom.MatchString(url):
		routers.ChatroomHandler(w, r)

	case websocketSever.MatchString(url):
		websocket.Handler(wsserver.WebSocketHandler).ServeHTTP(w, r)

	case publicFiles.MatchString(url):
		http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))).ServeHTTP(w, r)

	case adminLogin.MatchString(url):
		routers.AdminLoginHandler(w, r)

	case adminMain.MatchString(url):
		routers.AdminMainHandler(w, r)

	case adminUsers.MatchString(url):
		routers.AdminUsersHandler(w, r)

	case adminChatrooms.MatchString(url):
		routers.AdminChatroomsHandler(w, r)

	case adminUser.MatchString(url):
		routers.AdminUserHandler(w, r)

	case adminChatroom.MatchString(url):
		routers.AdminChatroomHandler(w, r)

	case adminLogout.MatchString(url):
		routers.AdminLogoutHandler(w, r)

	default:
		http.NotFound(w, r)
	}
}

func main() {
	mux := MyMux{}
	http.ListenAndServe(":80", mux)
}
