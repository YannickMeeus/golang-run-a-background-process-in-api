package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	go sayHiEvery5Seconds()

	port := "8000"
	http.HandleFunc("/_ping", ping)
	log.Printf("Server started at port %v", port)
	log.Fatal(http.ListenAndServe("localhost:" + port, nil))
}

func sayHiEvery5Seconds() {
	for {
		fmt.Println("Hello! ð")
		time.Sleep(5 * time.Second)
	}
}
