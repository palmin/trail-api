package main

import (
	"encoding/json"
	"os"
)

// Trail is a hiking trail as the Trailhead app shows it.
type Trail struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Kilometers float64 `json:"kilometers"`
	Climb      int     `json:"climb"`
	Difficulty string  `json:"difficulty"`
}

func loadTrails(path string) ([]Trail, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var trails []Trail
	if err := json.Unmarshal(data, &trails); err != nil {
		return nil, err
	}
	return trails, nil
}
