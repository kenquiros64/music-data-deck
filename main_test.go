package main

import (
	"io/ioutil"
	"music-data-deck/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zenazn/goji/web"
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
