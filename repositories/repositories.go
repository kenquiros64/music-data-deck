package repositories

import (
	"database/sql"
	"music-data-deck/models"
)

// GetAllSongs get all songs from database. No parameters. Return an array with the list of songs
func GetAllSongs() []models.Song {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()
	checkErr(err)

	// query
	rows, err := db.Query("SELECT s.id, s.artist, s.song, g.name, s.length FROM Songs s INNER JOIN Genres g ON s.genre = g.id")
	checkErr(err)
	var id int
	var artist string
	var song string
	var genre string
	var length int

	songs := []models.Song{}

	for rows.Next() {
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		checkErr(err)
		songs = append(songs,
			models.Song{
				ID:     id,
				Artist: artist,
				Song:   song,
				Genre:  genre,
				Length: length,
			})
	}

	rows.Close() //good habit to close

	db.Close()

	return songs
}

//Return an error when something happen
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
