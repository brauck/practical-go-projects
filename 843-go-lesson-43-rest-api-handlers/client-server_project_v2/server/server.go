package server

import (
	"log"
	"net/http"
)

type Server struct {
	storage *UserStorage
}

func NewServer() *Server {
	return &Server{
		storage: NewUserStorage("server/users.json"),
	}
}

func (s *Server) Start() {
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", s.Routes())
}
