package main

import (
	"fmt"
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

func (cfg *apiConfig) requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)

	data := fmt.Sprintf("Hits: %v", cfg.fileserverHits.Load())
	w.Write([]byte(data))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("200 OK"))
}

func main() {
	apiCfg := apiConfig{}

	filepathRoot := http.Dir(".")
	fileServer := http.StripPrefix("/app", http.FileServer(filepathRoot))
	port := "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fileServer))
	mux.HandleFunc("/healthz", healthzHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
