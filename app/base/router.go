package base

import (
	"fmt"
	"net/http"
)

// Router serves http
type Router struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

// NewRouter creates instance of Router
func NewRouter() *Router {
	router := new(Router)
	router.handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	return router
}

// ServeHTTP is called for every connection
func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		bad(w)
		return
	}
	f(w, r)
}

// Get sets get handler
func (s *Router) Get(path string, f http.HandlerFunc) {
	s.handlers[key("GET", path)] = f
}

// Post sets post handler
func (s *Router) Post(path string, f http.HandlerFunc) {
	s.handlers[key("POST", path)] = f
}

// Delete sets delete handler
func (s *Router) Delete(path string, f http.HandlerFunc) {
	s.handlers[key("DELETE", path)] = f
}

// Put sets put handler
func (s *Router) Put(path string, f http.HandlerFunc) {
	s.handlers[key("PUT", path)] = f
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"not found"}`))
}
