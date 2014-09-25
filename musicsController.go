package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMusics(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	getParams := r.Form
	//	urlParams := mux.Vars(r)

	query, qps := makeQuery("select * from music", getParams, map[string]string{"artist_id": "=", "title": "like"})
	log.Println("query: " + query)

	rows, err := Database.Query(query, qps...)
	if err != nil {
		RenderError(w, err)
	}
	defer rows.Close()

	musics := []Music{}

	for rows.Next() {
		music := Music{}
		err := rows.Scan(&music.Id, &music.ArtistId, &music.Title, &music.Outline)
		if err != nil {
			RenderError(w, err)
		} else {
			musics = append(musics, music)
		}
	}

	err = rows.Err()
	if err != nil {
		RenderError(w, err)
	}

	RenderJSON(w, musics)
}

func GetMusicId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	getParams := r.Form

	urlParams := mux.Vars(r)
	id := urlParams["id"]

	if id == "recent" {
		GetRecentMusics(w, r)
		return
	}

	getParams.Set("id", id)
	getParams.Set("limit", "1")
	query, qps := makeQuery("select * from music", getParams, map[string]string{"id": "="})

	js, _ := json.Marshal(qps)
	log.Println("query: " + query + " params: " + string(js))

	rows, err := Database.Query(query, qps...)
	if err != nil {
		Render404(w)
		return
	}
	defer rows.Close()

	musics := []Music{}

	for rows.Next() {
		music := Music{}
		err := rows.Scan(&music.Id, &music.ArtistId, &music.Title, &music.Outline)
		if err != nil {
			Render404(w)
		} else {
			musics = append(musics, music)
		}
	}

	err = rows.Err()
	if err != nil {
		Render404(w)
	}

	if len(musics) == 0 {
		Render404(w)
	}

	music := musics[0]

	RenderJSON(w, music)
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
	r.ParseForm()
	getParams := r.Form

	sql := "select * from play_history order by created_at desc"
	query, qps := makeQuery(sql, getParams, map[string]string{})

	rows, err := Database.Query(query, qps...)
	if err != nil {
		RenderError(w, err)
		return
	}
	defer rows.Close()

	histories := []PlayHistory{}

	for rows.Next() {
		history := PlayHistory{}
		err := rows.Scan(&history.Id, &history.MusicId, &history.CreatedAt)
		if err != nil {
			RenderError(w, err)
			return
		} else {
			histories = append(histories, history)
		}
	}

	err = rows.Err()
	if err != nil {
		RenderError(w, err)
		return
	}

	RenderJSON(w, histories)
}
