// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi gRPC client types
//
// Command:
// $ goa gen chat-api/design

package client

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
	chatapipb "chat-api/gen/grpc/chatapi/pb"

	goa "goa.design/goa/v3/pkg"
)

// NewGetchatRequest builds the gRPC request type from the payload of the
// "getchat" endpoint of the "chatapi" service.
func NewGetchatRequest(payload *chatapi.GetchatPayload) *chatapipb.GetchatRequest {
	message := &chatapipb.GetchatRequest{
		Id: int32(payload.ID),
	}
	return message
}

// NewGetchatResult builds the result type of the "getchat" endpoint of the
// "chatapi" service from the gRPC response type.
func NewGetchatResult(message *chatapipb.GoaChatCollection) chatapiviews.GoaChatCollectionView {
	result := make([]*chatapiviews.GoaChatView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &chatapiviews.GoaChatView{
			UserID:   &val.UserId,
			RoomName: &val.RoomName,
			Member:   &val.Member,
			Chat:     &val.Chat,
			PostDt:   &val.PostDt,
		}
		idptr := int(val.Id)
		result[i].ID = &idptr
	}
	return result
}

// NewPostchatRequest builds the gRPC request type from the payload of the
// "postchat" endpoint of the "chatapi" service.
func NewPostchatRequest(payload *chatapi.PostchatPayload) *chatapipb.PostchatRequest {
	message := &chatapipb.PostchatRequest{
		Id:       payload.ID,
		UserId:   payload.UserID,
		RoomName: payload.RoomName,
		Member:   payload.Member,
		Chat:     payload.Chat,
		PostDt:   payload.PostDt,
		Cookie:   payload.Cookie,
	}
	return message
}

// NewPostchatResult builds the result type of the "postchat" endpoint of the
// "chatapi" service from the gRPC response type.
func NewPostchatResult(message *chatapipb.PostchatResponse) bool {
	result := message.Field
	return result
}

// ValidateGoaChatCollection runs the validations defined on GoaChatCollection.
func ValidateGoaChatCollection(message *chatapipb.GoaChatCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateGoaChat(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateGoaChat runs the validations defined on GoaChat.
func ValidateGoaChat(message *chatapipb.GoaChat) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.PostDt", message.PostDt, goa.FormatDateTime))

	return
}
