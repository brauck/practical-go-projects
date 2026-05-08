package main

import (
	"fmt"
	"log"
	"net/http"
	"server/dowork"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] request received")
		// retrieve the context of the request
		rCtx := r.Context()
		// create the result channel
		resChan := make(chan int)
		// launch the function doWork in a goroutine
		go dowork.DoWork(rCtx, resChan)
		// Wait for
		// 1. the client drops the connection.
		// 2. the function doWork to finish it works
		select {
		case <-rCtx.Done():
			log.Println("[Handler] context canceled in main handler, client has diconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			fmt.Fprintf(w, "Response %d", result) // send data to client side
			return
		}
	})
	err := http.ListenAndServe("127.0.0.1:8989", nil) // set listen port
	if err != nil {
		panic(err)
	}
}
