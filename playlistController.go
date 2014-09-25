package main

import (
	"database/sql"
	"log"
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

func parseMusics(rows *sql.Rows) (musics []Music, plDetails []PlaylistDetail, theError error) {
	musics = []Music{}
	plDetails = []PlaylistDetail{}

	for rows.Next() {
		music := Music{}
		playlistDetail := PlaylistDetail{}
		err := rows.Scan(&music.Id, &music.ArtistId, &music.Title, &music.Outline, &playlistDetail.PlaylistName, &playlistDetail.Number, &playlistDetail.MusicId)
		if err != nil {
			theError = err
			return
		} else {
			musics = append(musics, music)
			plDetails = append(plDetails, playlistDetail)
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

	params := map[string]string{"name": "="}
	query, qps := makeQuery("select * from playlist", getParams, params)

	rows, err := Database.Query(query, qps...)
	if err != nil {
		Render404(w)
		log.Println(err)
		return
	}
	defer rows.Close()

	playlists, _ := parsePlaylists(rows)

	if len(playlists) == 0 {
		Render404(w)
		log.Println(err)
	}

	playlist := playlists[0]

	// get music

	//	getParams.Set("playlist_name", name)
	//	query, qps = makeQuery("select * from music inner join playlist_detail on music.id=playlist_detail.music_id where playlist_name = ? order by playlist_detail.number asc", getParams, queryParams)
	query = "select * from music inner join playlist_detail on music.id=playlist_detail.music_id where playlist_name = ? order by playlist_detail.number asc"
	rows, err = Database.Query(query, name)
	if err != nil {
		Render404(w)
		log.Println(err)
		return
	}
	defer rows.Close()

	musics, _, musErr := parseMusics(rows)

	data := map[string]interface{}{}
	data["name"] = playlist.Name
	data["outline"] = playlist.Outline
	data["musics"] = musics

	RenderJSON(w, data)
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
