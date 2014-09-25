package main

import "time"

type Artist struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Outline string `json:"outline"`
}

type Music struct {
	Id       uint64 `json:"id"`
	ArtistId uint64 `json:"artist_id"`
	Title    string `json:"title"`
	Outline  string `json:"outline"`
}

type Playlist struct {
	Name    string `json:"name"`
	Outline string `json:"outline"`
}

type PlaylistDetail struct {
	PlaylistName string `json:"playlist_name"`
	MusicId      uint64 `json:"music_id"`
	Number       uint64 `json:"number"`
}

type PlayHistory struct {
	Id        uint64    `json:"-"`
	MusicId   uint64    `json:"music_id"`
	CreatedAt time.Time `json:"last_played"`
}
