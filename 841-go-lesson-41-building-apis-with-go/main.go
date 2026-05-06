package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go API")
	io.WriteString(w, "Hello, world!\n")
	handler(w, r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := []User{
		{
			ID:   1,
			Name: "Alice",
		},
		{
			ID:   2,
			Name: "Bob",
		},
		{
			ID:   3,
			Name: "Jane",
		},
	}

	handler(w, r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request"))
	case http.MethodPost:
		w.Write([]byte("POST request"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	r.Method = "POST"
	var user []User
	json.NewDecoder(r.Body).Decode(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	fmt.Println(user)
	handler(w, r)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/create", createUser)

	//http.HandleFunc("/", handler)
	//http.HandleFunc("/user", handler)

	req := httptest.NewRequest("POST", "/user", nil)
	w := httptest.NewRecorder()

	userHandler(w, req)

	fmt.Println(w.Body.String())

	log.Println("run server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
