package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	trails, err := loadTrails("trails.json")
	if err != nil {
		log.Fatalf("no trails: %v", err)
	}

	mux := http.NewServeMux()
	routes(mux, trails)

	log.Printf("serving %d trails on port %s", len(trails), port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
