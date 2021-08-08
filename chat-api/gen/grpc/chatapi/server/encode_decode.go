// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi gRPC server encoders and decoders
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
	chatapipb "chat-api/gen/grpc/chatapi/pb"
	"context"
	"strings"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeGetchatResponse encodes responses from the "chatapi" service "getchat"
// endpoint.
func EncodeGetchatResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(chatapiviews.GoaChatCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatapi", "getchat", "chatapiviews.GoaChatCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewGoaChatCollection(result)
	return resp, nil
}

// DecodeGetchatRequest decodes requests sent to "chatapi" service "getchat"
// endpoint.
func DecodeGetchatRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		key string
		err error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			key = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chatapipb.GetchatRequest
		ok      bool
	)
	{
		if message, ok = v.(*chatapipb.GetchatRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chatapi", "getchat", "*chatapipb.GetchatRequest", v)
		}
	}
	var payload *chatapi.GetchatPayload
	{
		payload = NewGetchatPayload(message, key)
		if strings.Contains(payload.Key, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Key, " ", 2)[1]
			payload.Key = cred
		}
	}
	return payload, nil
}

// EncodePostchatResponse encodes responses from the "chatapi" service
// "postchat" endpoint.
func EncodePostchatResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(bool)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatapi", "postchat", "bool", v)
	}
	resp := NewPostchatResponse(result)
	return resp, nil
}

// DecodePostchatRequest decodes requests sent to "chatapi" service "postchat"
// endpoint.
func DecodePostchatRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		key string
		err error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			key = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chatapipb.PostchatRequest
		ok      bool
	)
	{
		if message, ok = v.(*chatapipb.PostchatRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chatapi", "postchat", "*chatapipb.PostchatRequest", v)
		}
		if err = ValidatePostchatRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *chatapi.PostchatPayload
	{
		payload = NewPostchatPayload(message, key)
		if strings.Contains(payload.Key, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Key, " ", 2)[1]
			payload.Key = cred
		}
	}
	return payload, nil
}
