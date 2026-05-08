package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	rootCtx := context.Background()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
	log.Println("req sent")
	if err != nil {
		panic(err)
	}
	// create context
	ctx, cancel := context.WithTimeout(rootCtx, 150*time.Millisecond)
	defer cancel()
	// attach context to our request
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp received", resp)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
	fmt.Println()
}
