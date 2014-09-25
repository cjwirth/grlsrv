package main

import "time"

type Artist struct {
	Id      uint64
	Name    string
	Outline string
}

type Music struct {
	Id       uint64
	ArtistId uint64
	Title    string
	Outline  string
}

type Playlist struct {
	Name    string
	Outline string
}

type PlaylistDetail struct {
	PlaylistName string
	MusicId      uint64
	Number       uint64
}

type PlayHistory struct {
	Id        uint64
	MusicId   uint64
	CreatedAt time.Time
}
