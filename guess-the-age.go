package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GuessTheAgeRequest struct {
	Name string `json:"name"`
}

type GuessTheAgeResponse struct {
	Age int `json:"age"`
}

type AgifyResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func handleGuessTheAgePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body GuessTheAgeRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pathTemplate := "https://api.agify.io/?name=%s"
		agifyResponse, err := http.Get(fmt.Sprintf(pathTemplate, body.Name))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		agifyResponseData, err := ioutil.ReadAll(agifyResponse.Body)

		var agifyResponseParsed AgifyResponse
		err = json.Unmarshal(agifyResponseData, &agifyResponseParsed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(GuessTheAgeResponse{Age: agifyResponseParsed.Age})
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
