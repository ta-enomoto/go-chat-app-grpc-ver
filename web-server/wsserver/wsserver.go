//WebSocketを実行するためのハンドラ
package wsserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/url"
	"strconv"
	"time"
	"unsafe"
	"web-server/query"
	"web-server/sessions"
)

//WebSocket用チャット構造体。通常のチャット構造体とは、①IdがString、②Cookieが追加されている点が異なる。
//①はJSONのデコード後のマッピングのため
//②はチャットの投稿者が誰なのか特定するため
type WsChat struct {
	Id       string    `json:"id"`
	UserId   string    `json:"userid"`
	RoomName string    `json:"roomname"`
	Member   string    `json:"member"`
	Chat     string    `json:"chat"`
	PostDt   time.Time `json:"postdt"`
	Cookie   string    `json:"cookie"`
}

//WebSocket用チャットルーム構造体。チャネル処理のためのルームで、通常のチャットルーム構造体とは別物
type WsChatroom struct {
	id      string
	forward chan string
	Join    chan *WsClient
	Leave   chan *WsClient
	clients map[*WsClient]bool
}

//メッセージの構造体
type Message struct {
	Message string
}

//WebSocketで管理するクライントの構造体
type WsClient struct {
	Send chan string
	Room *WsChatroom
}

//WebSocket用チャットルームのマネージャ
//WebSocket用チャットルームのマップWsRoomsでは、ルームIDをキーにチャットルームを管理する
type MANAGER struct {
	WsRooms map[string]*WsChatroom
}

//サーバー起動時にセッションマネージャも初期化
var Manager *MANAGER

func init() {
	Manager = NewManager()
}

//init()でセッションマネージャ初期化時に使う関数
func NewManager() *MANAGER {
	database := make(map[string]*WsChatroom)
	return &MANAGER{WsRooms: database}
}

func WebSocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	//JSONデータを受信、デコード後、WebSocket用チャット構造体にマッピングする
	var chatroomJson string
	if err := websocket.Message.Receive(ws, &chatroomJson); err == nil {
		var chatroom WsChat
		json.Unmarshal([]byte(chatroomJson), &chatroom)
		roomId := chatroom.Id

		//以下、アクセスしたいルームのWebSocket用チャットルームが既に存在する場合と、まだ存在しない場合で処理を分岐
		if _, exist := Manager.WsRooms[roomId]; exist {
			//既にWebSocket用チャットルームが存在していた場合の処理

			//WebSocket用クライアント構造体を初期化し、Roomフィールドに当該WebSocket用チャットルームを登録
			WsClient := &WsClient{
				Send: make(chan string),
				Room: Manager.WsRooms[roomId],
			}

			//WebSocket用チャットルームのJoinチャンネルにWebSocket用クライアントを送信
			Manager.WsRooms[roomId].Join <- WsClient

			//WebSocket用チャットルームのLeaveチャンネルにWebSocket用クライアントが送信され、
			//Ws用ルーム中のクライント数が0になったら、Ws用ルームを削除する
			defer func() {
				Manager.WsRooms[roomId].Leave <- WsClient
				if len(Manager.WsRooms[roomId].clients) == 0 {
					delete(Manager.WsRooms, roomId)
					fmt.Println("WebSocket用ルーム削除")
				}
			}()

			//メッセージ受信時の各クライアントへの書き出しを行う
			go WsClient.Write(ws)

			//各クライアントからのメッセージを受信する
			WsClient.Read(ws)
		} else {
			//まだアクセスしたルームのWebSocket用チャットルームが存在しない時の処理(Wsルーム生成)

			//WebSocket用チャットルーム構造体を初期化する
			WsChatroom := &WsChatroom{
				forward: make(chan string),
				Join:    make(chan *WsClient),
				Leave:   make(chan *WsClient),
				clients: make(map[*WsClient]bool),
			}

			//WebSocket用チャットルームを起動し、入退室・メッセージの送受信を監視する
			go WsChatroom.ChatroomRun()

			//WebSocket開通時に送信するメッセージを元に、ルームIDを設定する
			var chatroom WsChat
			json.Unmarshal([]byte(chatroomJson), &chatroom)
			WsChatroom.id = chatroom.Id

			//Ws用ルームマネージャに、生成したルームを登録する
			Manager.WsRooms[WsChatroom.id] = WsChatroom

			//WebSocket用クライアント構造体を初期化し、Roomフィールドに当該WebSocket用チャットルームを登録
			WsClient := &WsClient{
				Send: make(chan string),
				Room: WsChatroom,
			}

			//WebSocket用チャットルームのJoinチャンネルにWebSocket用クライアントを送信

			//WebSocket用チャットルームのLeaveチャンネルにWebSocket用クライアントが送信され、
			//Ws用ルーム中のクライント数が0になったら、Ws用ルームを削除する
			WsChatroom.Join <- WsClient
			defer func() {
				WsChatroom.Leave <- WsClient
				Manager.WsRooms[WsChatroom.id].Leave <- WsClient
				if len(Manager.WsRooms[WsChatroom.id].clients) == 0 {
					delete(Manager.WsRooms, WsChatroom.id)
					fmt.Println("WebSocket用ルーム削除")
				}
			}()

			//メッセージ受信時の各クライアントへの書き出しを行う
			go WsClient.Write(ws)

			//各クライアントからのメッセージを受信する
			WsClient.Read(ws)
		}
	}
}

//WebSocket用チャットルームへの入退室・メッセージの送受信を監視する関数
func (WsChatroom *WsChatroom) ChatroomRun() {
	for {
		select {
		case WsClient := <-WsChatroom.Join:
			//ルームへの入室処理

			WsChatroom.clients[WsClient] = true
			fmt.Printf("クライアントが入室しました。現在 %x 人のクライアントが存在します。\n", len(WsChatroom.clients))

		case WsClient := <-WsChatroom.Leave:
			//ルームからの退室処理

			//ルームからWs用クライアントを削除する
			delete(WsChatroom.clients, WsClient)
			fmt.Printf("クライアントが退出しました。現在 %x 人のクライアントが存在します。\n", len(WsChatroom.clients))

		case msg := <-WsChatroom.forward:
			//メッセージ受信時の処理
			fmt.Println("メッセージ受信")
			for target := range WsChatroom.clients {
				select {
				case target.Send <- msg:
					fmt.Println("メッセージ送信成功")
				default:
					fmt.Println("メッセージ送信失敗")
					delete(WsChatroom.clients, target)
				}
			}
		}
	}
}

//投稿されたメッセージを読み出すための関数
func (wc *WsClient) Read(ws *websocket.Conn) {

	for {
		//JSONデータを受信、デコード後、WebSocket用チャット構造体にマッピングする
		var msg string
		var receivedChat WsChat
		if err := websocket.Message.Receive(ws, &msg); err == nil {
			json.Unmarshal([]byte(msg), &receivedChat)
		} else {
			fmt.Println("json受信失敗")
			return
		}

		//チャットルームDBに接続する
		dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbChtrm.Close()

		//送信用に通常のチャット構造体を初期化
		sendingChat := new(query.Chat)

		//ルームIDを数値に変換し、DBからルーム情報を取得する
		roomId, _ := strconv.Atoi(receivedChat.Id)
		currentChatroom := query.SelectChatroomById(roomId, dbChtrm)

		sendingChat.Chatroom.Id = currentChatroom.Id
		sendingChat.Chatroom.RoomName = currentChatroom.RoomName

		//受信メッセージのcookieから、投稿者のセッション変数(ユーザーID)を取得する
		cookie, _ := url.QueryUnescape(receivedChat.Cookie)
		userSessionVar := session.Manager.SessionStore[cookie].SessionValue["userId"]

		//投稿者がルーム作成者と同じかどうかで、送信用チャットの投稿者を変える
		if userSessionVar == currentChatroom.UserId {
			//投稿主と部屋作成者が同じ場合
			sendingChat.Chatroom.UserId = currentChatroom.UserId
			sendingChat.Chatroom.Member = currentChatroom.Member
		} else {
			//投稿主と部屋作成者が違う場合(=メンバーが投稿主の場合)
			sendingChat.Chatroom.UserId = currentChatroom.Member
			sendingChat.Chatroom.Member = currentChatroom.UserId
		}

		sendingChat.Chat = receivedChat.Chat
		sendingChat.PostDt = time.Now().UTC().Round(time.Second)

		//送信用チャットをJSONにエンコード・byte型をstring型に変換し、Ws用ルームのforwardチャネルに送信する
		msgjson, err := json.Marshal(sendingChat)
		if err != nil {
			fmt.Println(err.Error())
		}
		msg = *(*string)(unsafe.Pointer(&msgjson))
		fmt.Println(msg)
		wc.Room.forward <- msg
		continue
	}
}

//受信したチャットを各クライアントのSendチャネルに送信する関数
func (wc *WsClient) Write(ws *websocket.Conn) {
	for msg := range wc.Send {
		if err := websocket.Message.Send(ws, msg); err != nil {
			break
		}
	}
}
