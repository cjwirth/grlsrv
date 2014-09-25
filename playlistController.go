package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func parsePlaylists(rows *sql.Rows) (playlists []Playlist, theError error) {
	playlists = []Playlist{}

	for rows.Next() {
		playlist := Playlist{}
		err := rows.Scan(&playlist.Name, &playlist.Outline)
		if err != nil {
			theError = err
			return
		} else {
			playlists = append(playlists, playlist)
		}
	}

	err := rows.Err()
	if err != nil {
		theError = err
		return
	}
	return
}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	getParams := r.Form
	urlParams := mux.Vars(r)
	name := urlParams["name"]
	getParams.Set("name", name)
	getParams.Set("limit", "1")

	params := map[string]string{"name": "=", "limit": "="}
	query, qps := makeQuery("select * from playlist", getParams, params)

	rows, err := Database.Query(query, qps...)
	if err != nil {
		Render404(w)
		return
	}
	defer rows.Close()

	playlists, err := parsePlaylists(rows)
	if err != nil {

	}

	if len(playlists) == 0 {
		Render404(w)
	}

	RenderJSON(w, playlists[0])
}

func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func AddPlaylistSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func RemovePlaylistSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
