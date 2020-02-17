package routes

import (
	"components"
	"net/http"
)

// Api: to handle routing
func Api() {
	http.HandleFunc("/", components.TestHandler())
}
