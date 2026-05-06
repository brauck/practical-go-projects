package server

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	storage *UserStorage
}

func NewServer() *Server {
	return &Server{
		storage: NewUserStorage("users.json"),
	}
}

func (s *Server) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getUsers(w, r)
	case http.MethodPost:
		s.createUser(w, r)
	case http.MethodPut:
		s.updateUser(w, r)
	case http.MethodDelete:
		s.deleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	users := s.storage.GetAll()
	json.NewEncoder(w).Encode(users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	user := s.storage.Create(input.Name)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	user, ok := s.storage.Update(input.ID, input.Name)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int `json:"id"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	ok := s.storage.Delete(input.ID)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
