package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()

	// Подключаем middleware
	r.Use(Logger)
	r.Use(RequestTimer)

	// Маршруты
	r.Get("/users", s.getUsers)
	r.Post("/users", s.createUser)
	r.Put("/users/{id}", s.updateUser)
	r.Delete("/users/{id}", s.deleteUser)

	return r
}
