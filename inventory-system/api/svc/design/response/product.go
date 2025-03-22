package response

import (
	. "goa.design/goa/v3/dsl"
)

var FindProductResult = ResultType("FindProductResult", func() {
	Field(1, "productId", Int)
	Field(2, "productName", String)
	Field(3, "productDescription", String)
	Field(4, "productMinStock", Int32)
})
