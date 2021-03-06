package main

import (
	"fmt"
	"music-data-deck/controllers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs"), controllers.AllSongs)
	mux.HandleFunc(pat.Get("/songs/search/:value"), controllers.SearchSong)
	mux.HandleFunc(pat.Get("/genres/search/:value"), controllers.SearchGenre)
	mux.HandleFunc(pat.Get("/genres/info"), controllers.GenresSongInfo)
	mux.HandleFunc(pat.Get("/songs/search/:min/:max"), controllers.SongsByLength)

	mux.Use(customLog)
	http.ListenAndServe("localhost:8000", mux)
}

func customLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %v\n", r.URL)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
