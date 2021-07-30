package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("getchat", func() {
	Title("Chat Service")
	Description("Service for chat app, a Goa teaser")
	Server("chat api", func() {
		Host("172.25.0.3", func() {
			URI("http://172.25.0.3:8000")
			//URI("grpc://172.25.0.2:8080")
		})
	})
})

var _ = Service("chatapi", func() {
	Description("The service performs get chat.")
	cors.Origin("http://172.25.0.2", func() {
		cors.Headers("Access-Control-Allow-Origin", "Authorization", "application/x-www-form-urlencoded")
		cors.Methods("GET")
		//cors.Expose("X-Time")
		//cors.MaxAge(600)
		cors.Credentials()
	})
	Method("getchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKey("api_key", "key", String, "API key used to perform authorization")
			Attribute("id", Int, func() {
				Description("room id")
				//Meta("rpc:tag", "1")
			})
			Required("key", "id")
		})
		Result(CollectionOf(Chat))
		Error("NotFound")
		Error("BadRequest")
		HTTP(func() {
			GET("/chatroom/{id}")
			Header("key:Authorization")
			Response(StatusOK)
		})
		//GRPC(func() {
		//	Response(CodeOK)
		//})
	})
	Method("postchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKey("api_key", "key", String, "API key used to perform authorization")
			Attribute("Id", String, "room id")
			Attribute("UserId", String, "user id")
			Attribute("RoomName", String, "room name")
			Attribute("Member", String, "member")
			Attribute("Chat", String, "chat")
			Attribute("PostDt", String, func() { Format(FormatDateTime) })
			Attribute("Cookie", String, "cookie")
			Required("key", "Id", "UserId", "RoomName", "Member", "Chat", "PostDt", "Cookie")
		})
		Result(Boolean)
		Error("NotFound")
		Error("BadRequest")
		HTTP(func() {
			POST("/chatroom/chat")
			Header("key:Authorization")
			Response(StatusOK)
		})
	})
})

var Chat = ResultType("application/vnd.goa.chat", func() {
	Description("All chat")
	Attributes(func() {
		Attribute("Id", Int, "room id")
		Attribute("UserId", String, "user id")
		Attribute("RoomName", String, "room name")
		Attribute("Member", String, "member")
		Attribute("Chat", String, "chat")
		Attribute("PostDt", String, func() { Format(FormatDateTime) })
		Required("Id", "UserId", "RoomName", "Member", "Chat", "PostDt")
	})
})

/*
gRPC用
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
*/

var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})
