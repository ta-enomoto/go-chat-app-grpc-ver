package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("getchat", func() {
	Title("Chat Service")
	Description("Service for chat app, a Goa teaser")
	Server("chat api", func() {
		Host("172.26.0.3", func() {
			URI("grpc://172.26.0.3:8080")
		})
	})
})

var _ = Service("chatapi", func() {
	Description("The service performs get chat.")
	cors.Origin("http://172.26.0.2", func() {
		cors.Headers("Access-Control-Allow-Origin", "Authorization")
		cors.Methods("POST")
		//cors.Expose("X-Time") APIのキャッシュ時使用
		//cors.MaxAge(600)
		cors.Credentials()
	})
	Method("getchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKeyField(1, "api_key", "key", String, "API key used to perform authorization")
			Field(2, "id", Int, "room id")
			Required("key", "id")
		})
		Result(CollectionOf(Chat))
		Error("NotFound")
		Error("BadRequest")
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("postchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKeyField(1, "api_key", "key", String, "API key used to perform authorization")
			Field(2, "Id", String, "room id")
			Field(3, "UserId", String, "user id")
			Field(4, "RoomName", String, "room name")
			Field(5, "Member", String, "member")
			Field(6, "Chat", String, "chat")
			Field(7, "PostDt", String, func() { Format(FormatDateTime) })
			Field(8, "Cookie", String, "cookie")
			Required("key", "Id", "UserId", "RoomName", "Member", "Chat", "PostDt", "Cookie")
		})
		Result(Boolean)
		Error("NotFound")
		Error("BadRequest")
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

//gRPC用
var Chat = ResultType("application/vnd.goa.chat", func() {
	Description("All chat")
	Attributes(func() {
		Field(1, "Id", Int, "room id")
		Field(2, "UserId", String, "user id")
		Field(3, "RoomName", String, "room name")
		Field(4, "Member", String, "member")
		Field(5, "Chat", String, "chat")
		Field(6, "PostDt", String, func() { Format(FormatDateTime) })
		Required("Id", "UserId", "RoomName", "Member", "Chat", "PostDt")
	})
})

var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})
