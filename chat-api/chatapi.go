package getchat

import (
	chatapi "chat-api/gen/chatapi"
	"context"
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
	s.logger.Print("chatapi.getchat")
	return
}
