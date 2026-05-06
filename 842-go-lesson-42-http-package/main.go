package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go API")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)

	userAgent := r.Header.Get("User-Agent")

	fmt.Fprintln(w, userAgent)

	fmt.Println(userAgent)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Write([]byte("Hello " + name))
}

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/", nameHandler) // http://localhost:8080/?name=John

	server.ListenAndServe()
}
