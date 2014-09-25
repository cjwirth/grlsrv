package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RenderJSON(w http.ResponseWriter, obj interface{}) {
	json, _ := json.Marshal(obj)
	w.Write(json)
}

func strLen(input string) int {
	return len([]rune(input))
}

func GetMusics(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	getParams := r.Form
	//	urlParams := mux.Vars(r)

	query := "select * from music "
	qps := []interface{}{}

	artist := getParams.Get("artist_id")
	if strLen(artist) > 0 {
		query += "where artist_id = ? "
		qps = append(qps, artist)
	}

	query += "limit ?"
	limit := getParams.Get("limit")
	if strLen(limit) > 0 {
		qps = append(qps, limit)
	} else {
		qps = append(qps, 100)
	}

	rows, err := Database.Query(query, qps...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	musics := []Music{}

	for rows.Next() {
		music := Music{}
		err := rows.Scan(&music.Id, &music.ArtistId, &music.Title, &music.Outline)
		if err != nil {
			log.Fatal(err)
		} else {
			musics = append(musics, music)
		}

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	RenderJSON(w, musics)
}

func GetMusicId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func PostMusics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func PutMusics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func DeleteMusics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func PlayMusicId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func GetMusicTimes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func GetRecentMusics(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
