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
