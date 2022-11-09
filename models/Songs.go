package models

import "github.com/jinzhu/gorm"

type Response struct {
	gorm.Model
	Song []Song `json:"results"`
}

type Song struct {
	gorm.Model
	IdSong   int     `json:"trackId"`
	SongName string  `json:"trackName"`
	Artist   string  `json:"artistName"`
	Duration int     `json:"trackTimeMillis"`
	Album    string  `json:"collectionName"`
	Artwork  string  `json:"artworkUrl30"`
	Price    float32 `json:"trackPrice"`
}
