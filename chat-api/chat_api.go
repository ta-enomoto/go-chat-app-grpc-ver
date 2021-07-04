package getchat

import (
	"chat-api/db_info/query"
	chatapi "chat-api/gen/chatapi"
	"context"
	"database/sql"
	"fmt"
	"log"
)

// chat api service example implementation.
// The example methods log the requests and return zero values.
type chatAPIsrvc struct {
	logger *log.Logger
}

// NewChatAPI returns the chat api service implementation.
func NewChatAPI(logger *log.Logger) chatapi.Service {
	return &chatAPIsrvc{logger}
}

// GetChat implements get chat.
func (s *chatAPIsrvc) Getchat(ctx context.Context, p *chatapi.GetchatPayload) (res chatapi.GoaChatCollection, err error) {

	dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbChtrm.Close()

	selectedChatroom := query.SelectChatroomById(p.ID, dbChtrm)
	//userId := selectedChatroom.UserId
	//member := selectedChatroom.Member
	fmt.Println(selectedChatroom.Id)

	Chats := query.SelectAllChatsById(selectedChatroom.Id, dbChtrm)

	fmt.Println(p.ID)
	fmt.Println(Chats)
	fmt.Println("successed")
	s.logger.Print("chatAPI.get chat")
	return Chats, nil
}
