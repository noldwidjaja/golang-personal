package components

import (
	"fmt"
	"net/http"
)

// TestHandler : for testing purposes
func TestHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "test")
}
