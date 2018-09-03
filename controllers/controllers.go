package controllers

import (
	"encoding/json"
	"fmt"
	"music-data-deck/repositories"
	"net/http"
	"strconv"

	"goji.io/pat"
)

// AllSongs GET Method ["/songs"]
// Get all songs from the database
func AllSongs(w http.ResponseWriter, r *http.Request) {
	songs := repositories.GetAllSongs()
	if songs == nil {
		http.Error(w, "Something has occured. ", http.StatusInternalServerError)
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
		http.Error(w, "Something has occured. ", http.StatusInternalServerError)
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
		http.Error(w, "Something has occured. ", http.StatusInternalServerError)
	} else {
		jsonOut, _ := json.Marshal(genres)
		fmt.Fprintf(w, string(jsonOut))
	}
}

// GenresSongInfo GET Method ["/genres/info"]
// Return a list of the genres, and the number of songs and the total length of all the songs by genre.
func GenresSongInfo(w http.ResponseWriter, r *http.Request) {
	genresSongInfo := repositories.GenresSongInfo()

	if genresSongInfo == nil {
		http.Error(w, "Something has occured. ", http.StatusInternalServerError)
	} else {
		jsonOut, _ := json.Marshal(genresSongInfo)
		fmt.Fprintf(w, string(jsonOut))
	}
}

// SongsByLength GET Method ["/songs/search/:min/:max"]
// Return a list of the songs by passing a minimum and maximum length.
func SongsByLength(w http.ResponseWriter, r *http.Request) {
	min := pat.Param(r, "min")
	max := pat.Param(r, "max")

	// string to int
	minLength, errMin := strconv.Atoi(min)
	maxLength, errMax := strconv.Atoi(max)
	if errMin != nil || errMax != nil {
		http.Error(w, "Incorrect minimum or maximun value.", http.StatusBadRequest)
	}

	if minLength <= maxLength {
		songs := repositories.SongsByLength(minLength, maxLength)

		if songs == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			jsonOut, _ := json.Marshal(songs)
			fmt.Fprintf(w, string(jsonOut))
		}
	} else {
		http.Error(w, "Minimum length is greater than maximum.", http.StatusBadRequest)
	}
}
