package base

import "net/http"

// Server creates golang server
type Server struct {
	Router *Router
	port   string
}

// NewServer constructs the server with the router inside the server
func NewServer(portNumber string) *Server {
	s := &Server{
		Router: NewRouter(),
		port:   ":" + portNumber,
	}
	return s
}

// Serve serves the server for development
func (s *Server) Serve() error {
	return http.ListenAndServe(s.port, s.Router)
}

// Routes is where routing will be inserted
func (s *Server) Routes() {}
