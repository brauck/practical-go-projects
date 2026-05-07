package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var input struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	user, ok := s.storage.Update(id, input.Name)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	ok := s.storage.Delete(id)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
