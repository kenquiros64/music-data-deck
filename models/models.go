package models

// Song model structure
type Song struct {
	ID     int    `json:"id"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
}

// Genre model structure
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GenreSongInfo model structure
type GenreSongInfo struct {
	Name        string `json:"name"`
	TotalSongs  int    `json:"songs"`
	TotalLength int    `json:"length"`
}
