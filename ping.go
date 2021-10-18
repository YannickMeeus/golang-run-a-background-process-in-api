package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type pingResponse struct {
	Data string `json:"data"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		response := pingResponse{Data: "pong"}
		jsonResp, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		_, _ = w.Write(jsonResp)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
