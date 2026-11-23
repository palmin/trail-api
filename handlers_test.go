package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTrailsByDifficulty(t *testing.T) {
	trails, err := loadTrails("trails.json")
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	routes(mux, trails)

	response := httptest.NewRecorder()
	mux.ServeHTTP(response, httptest.NewRequest("GET", "/trails?difficulty=hard", nil))

	var found []Trail
	if err := json.Unmarshal(response.Body.Bytes(), &found); err != nil {
		t.Fatal(err)
	}
	if len(found) == 0 {
		t.Fatal("no hard trails")
	}
	for _, trail := range found {
		if trail.Difficulty != "hard" {
			t.Errorf("%s is %s", trail.Name, trail.Difficulty)
		}
	}
}
