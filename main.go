package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	directory := http.Dir(".")
	mux.Handle("/", http.FileServer(directory))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
