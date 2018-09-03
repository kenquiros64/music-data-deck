package repositories

import (
	"database/sql"
	"fmt"
	"music-data-deck/models"
)

// GetAllSongs get all songs from database. No parameters. Return an array with the list of songs
func GetAllSongs() []models.Song {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	if err != nil {
		return nil
	}
	// query
	rows, err := db.Query("SELECT s.id, s.artist, s.song, g.name, s.length FROM Songs s INNER JOIN Genres g ON s.genre = g.id")

	if err != nil {
		return nil
	}

	var id int
	var artist string
	var song string
	var genre string
	var length int

	songs := []models.Song{}

	for rows.Next() {
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		if err != nil {
			return nil
		}
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

// SearchSong get all songs from database when it matches with the given value. Receive string value. Return an array with the list of songs
func SearchSong(value string) []models.Song {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}
	// query
	rows, err := db.Query("SELECT s.id, s.artist, s.song, g.name, s.length FROM Songs s INNER JOIN Genres g ON s.genre = g.id WHERE s.artist = ? or s.song = ? or g.name = ?", value, value, value)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var id int
	var artist string
	var song string
	var genre string
	var length int

	songs := []models.Song{}

	for rows.Next() {
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		if err != nil {
			fmt.Println(err)
			return nil
		}
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

// SearchGenre get all genres from database when it matches with the given value. Receive string value. Return an array with the list of genres
func SearchGenre(value string) []models.Genre {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	if err != nil {
		return nil
	}
	// query
	rows, err := db.Query("SELECT g.id, g.name FROM Genres g WHERE g.name like ?", value)

	if err != nil {
		return nil
	}

	var id int
	var genre string

	genres := []models.Genre{}

	for rows.Next() {
		err = rows.Scan(&id, &genre)
		if err != nil {
			return nil
		}
		genres = append(genres,
			models.Genre{
				ID:   id,
				Name: genre,
			})
	}

	rows.Close() //good habit to close

	db.Close()

	return genres
}
