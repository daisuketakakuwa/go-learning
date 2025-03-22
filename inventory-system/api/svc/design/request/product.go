package request

import (
	. "goa.design/goa/v3/dsl"
)

var CreateProductPayload = Type("CreateProductPayload", func() {
	Field(1, "productName", String)
	Field(2, "productDescription", String)
	Field(3, "productMinStock", Int32)
	Required("productName")
})

var UpdateProductPayload = Type("UpdateProductPayload", func() {
	Field(1, "productId", Int)
	Field(2, "productName", String)
	Field(3, "productDescription", String)
	Field(4, "productMinStock", Int32)
	Required("productId")
})
