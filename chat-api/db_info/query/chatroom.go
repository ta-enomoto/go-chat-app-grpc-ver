//チャット情報を扱うクエリパッケージ
package query

import (
	"chat-api/db_info/config"
	chatapi "chat-api/gen/chatapi"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Chatroom struct {
	Id       int    `json:"id"`
	UserId   string `json:"userId"`
	RoomName string `json:"roomName"`
	Member   string `json:"member"`
}

var ConStrChtrm string

func init() {
	confDbChtrm, err := config.ReadConfDbChtrm()
	if err != nil {
		fmt.Println(err.Error())
	}
	ConStrChtrm = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=%s", confDbChtrm.User, confDbChtrm.Pass, confDbChtrm.Host, confDbChtrm.Port, confDbChtrm.DbName, confDbChtrm.Charset)
}

//idでチャットルームを選択
func SelectChatroomById(id int, db *sql.DB) (chatroom Chatroom) {

	err := db.QueryRow("SELECT * FROM ROOM_STRUCTS_OF_CHAT WHERE ID = ?", id).Scan(&chatroom.Id, &chatroom.UserId, &chatroom.RoomName, &chatroom.Member)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

//特定のチャットルームのチャットをすべて取得する
func SelectAllChatsById(id int, db *sql.DB) (chats []*chatapi.GoaChat) {

	rows, err := db.Query("SELECT * FROM ALL_STRUCTS_OF_CHAT WHERE ID = ?", id)
	if err != nil {
		return chats
	}

	for rows.Next() {
		chat := &chatapi.GoaChat{}
		err := rows.Scan(&chat.ID, &chat.UserID, &chat.RoomName, &chat.Member, &chat.Chat, &chat.PostDt)
		if err != nil {
			fmt.Println(err.Error())
		}
		chats = append(chats, chat)
	}
	return
}
