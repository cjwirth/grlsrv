package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/musics", GetMusics).Methods("GET")
	router.HandleFunc("/api/musics/:id", GetMusicId).Methods("GET")
	router.HandleFunc("/api/musics", GetMusics).Methods("POST")
	router.HandleFunc("/api/musics/:id", PutMusics).Methods("PUT")
	router.HandleFunc("/api/musics/:id", DeleteMusics).Methods("DELETE")
	router.HandleFunc("/api/musics/:id/play", PlayMusicId).Methods("POST")
	router.HandleFunc("/api/musics/times", GetMusicTimes).Methods("GET")
	router.HandleFunc("/api/musics/recent", GetRecentMusics).Methods("GET")

	router.HandleFunc("/api/playlists/:name", GetPlaylist).Methods("GET")
	router.HandleFunc("/api/playlists/:name", CreatePlaylist).Methods("POST")
	router.HandleFunc("/api/playlists/:name/add", AddPlaylistSong).Methods("POST")
	router.HandleFunc("/api/playlists/:name/remove", RemovePlaylistSong).Methods("POST")

	http.Handle("/", router)

	http.ListenAndServe(":9000", nil)
}
