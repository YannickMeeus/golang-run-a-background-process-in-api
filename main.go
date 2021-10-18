package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8000"
	http.HandleFunc("/_ping", ping)
	http.HandleFunc("/guess-the-age", handleGuessTheAgePost)
	log.Printf("Server started at port %v", port)
	log.Fatal(http.ListenAndServe("localhost:" + port, nil))
}
