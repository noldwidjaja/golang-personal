package main

import "github.com/noldwidjaja/golang-personal.git/app/components/test"

func (s *server) routes() {
	s.getRouter().Get("/test", test.HandleTest())
}
