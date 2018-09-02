package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"goji.io"
	"goji.io/pat"
)

const (
	dbhost = "https://s3.amazonaws.com/bv-challenge/jrdd.db"
	dbname = "jrdd.db"
)

func main() {
	initDb()
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func initDb() {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()
	checkErr(err)

	// query
	rows, err := db.Query("SELECT id, artist, song FROM Songs")
	checkErr(err)
	var id int
	var artist string
	var song string
	var genre int
	var length int

	for rows.Next() {
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(artist)
		fmt.Println(song)
		fmt.Println(genre)
		fmt.Println(length)
	}

	rows.Close() //good habit to close

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
