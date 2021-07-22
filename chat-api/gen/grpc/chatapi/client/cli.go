// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi gRPC client CLI support package
//
// Command:
// $ goa gen chat-api/design

package client

import (
	chatapi "chat-api/gen/chatapi"
	chatapipb "chat-api/gen/grpc/chatapi/pb"
	"encoding/json"
	"fmt"
)

// BuildGetchatPayload builds the payload for the chatapi getchat endpoint from
// CLI flags.
func BuildGetchatPayload(chatapiGetchatMessage string) (*chatapi.GetchatPayload, error) {
	var err error
	var message chatapipb.GetchatRequest
	{
		if chatapiGetchatMessage != "" {
			err = json.Unmarshal([]byte(chatapiGetchatMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 483739648181357268\n   }'")
			}
		}
	}
	v := &chatapi.GetchatPayload{
		ID: int(message.Id),
	}

	return v, nil
}
