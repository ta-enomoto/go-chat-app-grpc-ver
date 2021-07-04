// Code generated by goa v3.3.1, DO NOT EDIT.
//
// chatapi HTTP server types
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
)

// GoaChatResponseCollection is the type of the "chatapi" service "getchat"
// endpoint HTTP response body.
type GoaChatResponseCollection []*GoaChatResponse

// GoaChatResponse is used to define fields on response body types.
type GoaChatResponse struct {
	// room id
	ID int `form:"id" json:"id" xml:"id"`
	// user id
	UserID string `form:"UserId" json:"UserId" xml:"UserId"`
	// room name
	RoomName string `form:"RoomName" json:"RoomName" xml:"RoomName"`
	// member
	Member string `form:"Member" json:"Member" xml:"Member"`
	// chat
	Chat   string `form:"Chat" json:"Chat" xml:"Chat"`
	PostDt string `form:"PostDt" json:"PostDt" xml:"PostDt"`
}

// NewGoaChatResponseCollection builds the HTTP response body from the result
// of the "getchat" endpoint of the "chatapi" service.
func NewGoaChatResponseCollection(res chatapiviews.GoaChatCollectionView) GoaChatResponseCollection {
	body := make([]*GoaChatResponse, len(res))
	for i, val := range res {
		body[i] = marshalChatapiviewsGoaChatViewToGoaChatResponse(val)
	}
	return body
}

// NewGetchatPayload builds a chatapi service getchat endpoint payload.
func NewGetchatPayload(id int) *chatapi.GetchatPayload {
	v := &chatapi.GetchatPayload{}
	v.ID = id

	return v
}
