package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/_ping", ping)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}