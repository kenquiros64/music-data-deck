package main

import (
	"io/ioutil"
	"music-data-deck/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zenazn/goji/web"
	goji "goji.io"
	"goji.io/pat"
)

// ParseResponse parse response message
func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

// TestAllSongs Test songs
func TestAllSongs(t *testing.T) {
	m := web.New()
	m.Get("/songs", controllers.AllSongs)
	ts := httptest.NewServer(m)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/songs")
	if err != nil {
		t.Error("unexpected")
	}
	_, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("Invalid status code")
	}
}

// TestGenreInfo Test genre info
func TestGenreInfo(t *testing.T) {
	m := web.New()
	m.Get("/genres/info", controllers.GenresSongInfo)
	ts := httptest.NewServer(m)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/genres/info")
	if err != nil {
		t.Error("unexpected")
	}
	_, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("Invalid status code")
	}
}

// TestSearchSong Test for search songs
func TestSearchSong(t *testing.T) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs/search/:value"), controllers.SearchSong)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/songs/search/424")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("Invalid status code")
	}
	if c != `[{"artist":"424","song":"Gala","genre":"Indie Rock","length":189}]` {
		t.Error("Expected response doesn't match")
	}
}

// TestSearchGenre Test for search genres
func TestSearchGenre(t *testing.T) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/genres/search/:value"), controllers.SearchGenre)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/genres/search/rap")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("Invalid status code")
	}
	if c != `[{"name":"Rap"}]` {
		t.Error("Expected response doesn't match")
	}
}

// TestSongsByLength Test for search songs by length
func TestSongsByLength(t *testing.T) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs/search/:min/:max"), controllers.SongsByLength)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/songs/search/240/245")
	if err != nil {
		t.Error("unexpected")
	}
	c, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("Invalid status code")
	}
	if c != `[{"artist":"Bobby Darin","song":"Mack the Knife","genre":"Rock","length":245},{"artist":"Debby Boone","song":"You Light Up My Life","genre":"Pop","length":245}]` {
		t.Error("Expected response doesn't match")
	}
}

// TestSongsByLengthBadRequest Test for search songs by length
func TestSongsByLengthBadRequest(t *testing.T) {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs/search/:min/:max"), controllers.SongsByLength)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/songs/search/245/240")
	if err != nil {
		t.Error("unexpected")
	}
	_, s := ParseResponse(res)
	if s != http.StatusBadRequest {
		t.Error("Invalid status code")
	}
}
