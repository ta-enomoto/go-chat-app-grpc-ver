// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi gRPC server types
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
	chatapipb "chat-api/gen/grpc/chatapi/pb"
)

// NewGetchatPayload builds the payload of the "getchat" endpoint of the
// "chatapi" service from the gRPC request type.
func NewGetchatPayload(message *chatapipb.GetchatRequest) *chatapi.GetchatPayload {
	v := &chatapi.GetchatPayload{
		ID: int(message.Id),
	}
	return v
}

// NewGoaChatCollection builds the gRPC response type from the result of the
// "getchat" endpoint of the "chatapi" service.
func NewGoaChatCollection(result chatapiviews.GoaChatCollectionView) *chatapipb.GoaChatCollection {
	message := &chatapipb.GoaChatCollection{}
	message.Field = make([]*chatapipb.GoaChat, len(result))
	for i, val := range result {
		message.Field[i] = &chatapipb.GoaChat{}
		if val.ID != nil {
			message.Field[i].Id = int32(*val.ID)
		}
		if val.UserID != nil {
			message.Field[i].UserId = *val.UserID
		}
		if val.RoomName != nil {
			message.Field[i].RoomName = *val.RoomName
		}
		if val.Member != nil {
			message.Field[i].Member = *val.Member
		}
		if val.Chat != nil {
			message.Field[i].Chat = *val.Chat
		}
		if val.PostDt != nil {
			message.Field[i].PostDt = *val.PostDt
		}
	}
	return message
}