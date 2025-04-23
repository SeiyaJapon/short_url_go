package main

import (
	http "URL_shortener/Internal/infrastructure"
	"log"
)

func main() {
	server, err := http.NewServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	log.Println("Server started on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
