package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseBody := errorResponse{
		Error: msg,
	}

	data, err := json.Marshal(responseBody)
	if err != nil {
		log.Printf("erorr marshalling the json: %v\n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling the payload: %v\n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func cleanProfanity(chirp string) string {
	words := strings.Split(chirp, " ")
	profanity := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	for index, word := range words {
		if profanity[strings.ToLower(word)] {
			words[index] = "****"
		}
	}

	return strings.Join(words, " ")
}
