package main

import (
	"fmt"
	"log"
	"time"

	"rest_api_project/client"
	"rest_api_project/server"
)

func main() {
	// Запускаем сервер в отдельной горутине
	srv := server.NewServer()
	go srv.Start()

	// Даем серверу запуститься
	time.Sleep(300 * time.Millisecond)

	// Создаем клиент
	api := client.New("http://localhost:8080")

	// POST
	api.CreateUser(client.User{Name: "Jade"})
	api.CreateUser(client.User{Name: "Doe"})
	created, _ := api.CreateUser(client.User{Name: "Bob"})
	fmt.Println("Created:", created)
	api.CreateUser(client.User{Name: "John"})

	// GET
	users, _ := api.GetUsers()
	fmt.Println("Users now:", users)

	// PUT
	created.Name = "Alice Updated"
	updated, _ := api.UpdateUser(*created)
	fmt.Println("Updated:", updated)

	// GET
	users, _ = api.GetUsers()
	fmt.Println("Users now:", users)

	// DELETE
	if err := api.DeleteUser(created.ID); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user", created.ID)
}
