package test

import (
	"net/http"

	"github.com/noldwidjaja/golang-personal.git/app/base"
)

// HandleTest : test
func HandleTest() http.HandlerFunc {
	test := Test{1, "TestHandler v2"}
	res := base.NewResponse(test, "ok", 200)
	return res.HandleJSONResponse()
}
