package main

import (
	"encoding/json"
	"net/http"
)

func routes(mux *http.ServeMux, trails []Trail) {
	mux.HandleFunc("GET /trails", func(w http.ResponseWriter, r *http.Request) {
		listTrails(w, r, trails)
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})
}

// listTrails answers with every trail, or only those of the difficulty asked for.
func listTrails(w http.ResponseWriter, r *http.Request, trails []Trail) {
	difficulty := r.URL.Query().Get("difficulty")

	found := []Trail{}
	for _, trail := range trails {
		if difficulty == "" || trail.Difficulty == difficulty {
			found = append(found, trail)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(found)
}
