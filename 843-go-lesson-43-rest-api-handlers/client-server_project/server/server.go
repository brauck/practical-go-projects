package server

import (
	"log"
	"net/http"
)

func (s *Server) Start() {
	http.HandleFunc("/users", s.UsersHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
