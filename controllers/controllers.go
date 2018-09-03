package controllers

import (
	"encoding/json"
	"fmt"
	"music-data-deck/repositories"
	"net/http"

	"goji.io/pat"
)

// AllSongs GET Method ["/songs"]
// Get all songs from the database
func AllSongs(w http.ResponseWriter, r *http.Request) {
	songs := repositories.GetAllSongs()
	if songs == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		jsonOut, _ := json.Marshal(songs)
		fmt.Fprintf(w, string(jsonOut))
	}
}

// SearchSong GET Method ["/songs/search/:value"]
// Search songs by artist, song or genre
func SearchSong(w http.ResponseWriter, r *http.Request) {
	value := pat.Param(r, "value")
	songs := repositories.SearchSong(value)
	if songs == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		jsonOut, _ := json.Marshal(songs)
		fmt.Fprintf(w, string(jsonOut))
	}
}

// SearchGenre GET Method ["/genres/search/:value"]
// Search genres by name
func SearchGenre(w http.ResponseWriter, r *http.Request) {
	value := pat.Param(r, "value")
	genres := repositories.SearchGenre(value)

	if genres == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		jsonOut, _ := json.Marshal(genres)
		fmt.Fprintf(w, string(jsonOut))
	}
}
