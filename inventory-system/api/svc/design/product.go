package design

import (
	"inventory-system/api/svc/design/request"
	"inventory-system/api/svc/design/response"

	. "goa.design/goa/v3/dsl"
)

// APIサーバ定義
var _ = API("inventory-system", func() {
    // API の説明（タイトルと説明）
    Title("Inteventory System Service")
    Description("Service for adding numbers, a Goa teaser")

    // ホスト情報
    Server("inventory-system", func() {
        Host("localhost", func() {
            URI("http://localhost:8008") // HTTP REST API
            URI("grpc://localhost:8088") // gRPC
        })
    })
})

// サービス定義
var _ = Service("inventory", func() {
    Description("The calc service performs operations on numbers.")

    // Create
    Method("create", func() {
        Payload(request.CreateProductPayload)
        Result(String)
        HTTP(func() {
            POST("/products")
        })
        GRPC(func() {
            Response(CodeOK)
        })
    })
    // Update
    Method("update", func() {
        Payload(request.UpdateProductPayload)
        Result(String)
        HTTP(func() {
            PUT("/products/{productId}")
        })
        GRPC(func() {
            Response(CodeOK)
        })
    })
    // Find
    Method("find", func() {
        Payload(func(){
            Field(1, "productId", Int)
            Required("productId")
        })
        Result(response.FindProductResult)
        HTTP(func() {
            GET("/products/{productId}")
        })
        GRPC(func() {
            Response(CodeOK)
        })
    })
    // Delete
    Method("delete", func() {
        Payload(func(){
            Field(1, "productId", Int)
            Required("productId")
        })
        Result(String)
        HTTP(func() {
            DELETE("/products/{productId}")
        })
        GRPC(func() {
            Response(CodeOK)
        })
    })

})