package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// API 定義
var _ = API("getchat", func() {
	// API の説明（タイトルと説明）
	Title("Chat Service")
	Description("Service for chat app, a Goa teaser")
	// サーバ定義
	Server("chat api", func() {
		Host("172.25.0.4", func() {
			URI("http://172.25.0.4:8000") // HTTP REST API
			URI("grpc://172.25.0.4:8080") // gRPC
		})
	})
})

// サービス定義
var _ = Service("chatapi", func() {
	// 説明
	Description("The service performs get chat.")
	// メソッド (HTTPでいうところのエンドポントに相当)
	cors.Origin("http://172.25.0.2", func() {
		cors.Headers("Access-Control-Allow-Origin")
		cors.Methods("GET")
		cors.Expose("X-Time")
		cors.MaxAge(600)
		cors.Credentials()
	})
	Method("getchat", func() {
		// ペイロード定義
		Payload(func() {
			// 整数型の属性 `a` これは左の被演算子
			Attribute("id", Int, func() {
				Description("room id") // 説明
				Meta("rpc:tag", "1")   // gRPC 用のメタ情報。タグ定義
			})
			Required("id") // a と b は required な属性であることの指定
		})
		Result(CollectionOf(Chat)) // メソッドの返値（整数を返す）
		Error("NotFound")
		Error("BadRequest")
		// HTTP トランスポート用の定義
		HTTP(func() {
			GET("/chatroom/{id}") // GET エンドポイント
			Response(StatusOK)    // レスポンスのステータスは Status OK = 200 を返す
		})
		// GRPC トランスポート用の定義
		GRPC(func() {
			Response(CodeOK) // レスポンスのステータスは CodeOK を返す
		})
	})
})
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
