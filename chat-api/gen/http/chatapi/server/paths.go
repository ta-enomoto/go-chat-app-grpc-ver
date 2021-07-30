// Code generated by goa v3.4.3, DO NOT EDIT.
//
// HTTP request path constructors for the chatapi service.
//
// Command:
// $ goa gen chat-api/design

package server

import (
	"fmt"
)

// GetchatChatapiPath returns the URL path to the chatapi service getchat HTTP endpoint.
func GetchatChatapiPath(id int) string {
	return fmt.Sprintf("/chatroom/%v", id)
}

// PostchatChatapiPath returns the URL path to the chatapi service postchat HTTP endpoint.
func PostchatChatapiPath() string {
	return "/chatroom/chat"
}
