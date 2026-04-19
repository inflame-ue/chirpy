package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("200 OK"))
}

func main() {
	filepathRoot := http.Dir(".")
	fileServer := http.FileServer(filepathRoot)
	port := "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))
	mux.HandleFunc("/healthz", healthzHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
