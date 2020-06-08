package main

import (
	"github.com/noldwidjaja/golang-personal.git/app/base"
)

type server struct {
	Server *base.Server
}

func createServer() *server {
	s := &server{
		base.NewServer("4000"),
	}

	return s
}

func (s *server) serve() {
	s.Server.Serve()
}

func (s *server) getRouter() *base.Router {
	return s.Server.Router
}
