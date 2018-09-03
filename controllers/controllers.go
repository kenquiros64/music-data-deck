package controllers

import (
	"encoding/json"
	"fmt"
	"music-data-deck/repositories"
	"net/http"
)

// AllSongs GET Method ["/songs"]
// Get all songs from the database
func AllSongs(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(repositories.GetAllSongs())
	fmt.Fprintf(w, string(jsonOut))
}
