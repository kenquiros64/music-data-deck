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

	rows.Close() // close rows

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
	rows, err := db.Query("SELECT s.id, s.artist, s.song, g.name, s.length FROM Songs s INNER JOIN Genres g ON s.genre = g.id WHERE s.artist like ? or s.song like ? or g.name like ?", value, value, value)

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

	rows.Close() // close rows

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

	rows.Close() // close rows

	db.Close()

	return genres
}

// GenresSongInfo Return a list of the genres, and the number of songs and the total length of all the songs by genre.
// No parameters. Return an array with the list of genres
func GenresSongInfo() []models.GenreSongInfo {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	if err != nil {
		return nil
	}
	// query
	rows, err := db.Query("SELECT g.name, COUNT(s.song), SUM(s.length) FROM Songs s INNER JOIN Genres g ON s.genre = g.id GROUP BY g.name")

	if err != nil {
		return nil
	}

	var genre string
	var totalSongs int
	var totalLength int

	genres := []models.GenreSongInfo{}

	for rows.Next() {
		err = rows.Scan(&genre, &totalSongs, &totalLength)
		if err != nil {
			return nil
		}
		genres = append(genres,
			models.GenreSongInfo{
				Name:        genre,
				TotalSongs:  totalSongs,
				TotalLength: totalLength,
			})
	}

	rows.Close() // close rows

	db.Close()

	return genres
}

// SongsByLength get all songs between minimun and maximum length
func SongsByLength(minLength int, maxLength int) []models.Song {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}
	// query
	rows, err := db.Query("SELECT s.id, s.artist, s.song, g.name, s.length FROM Songs s INNER JOIN Genres g ON s.genre = g.id WHERE s.length between ? and ?", minLength, maxLength)

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

	rows.Close() // close rows

	db.Close()

	return songs

}
