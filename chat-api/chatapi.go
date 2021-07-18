package getchat

import (
	"chat-api/db_info/query"
	chatapi "chat-api/gen/chatapi"
	"context"
	"database/sql"
	"fmt"
	"log"
)

// chatapi service example implementation.
// The example methods log the requests and return zero values.
type chatapisrvc struct {
	logger *log.Logger
}

// NewChatapi returns the chatapi service implementation.
func NewChatapi(logger *log.Logger) chatapi.Service {
	return &chatapisrvc{logger}
}

// Getchat implements getchat.
func (s *chatapisrvc) Getchat(ctx context.Context, p *chatapi.GetchatPayload) (res chatapi.GoaChatCollection, err error) {

	dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbChtrm.Close()

	selectedChatroom := query.SelectChatroomById(p.ID, dbChtrm)
	fmt.Println(selectedChatroom.Id)

	Chats := query.SelectAllChatsById(selectedChatroom.Id, dbChtrm)
	fmt.Println(p.ID)
	fmt.Println("successed")
	s.logger.Print("chatAPI.get chat")
	fmt.Println(Chats)

	return Chats, nil
}
