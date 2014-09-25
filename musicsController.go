package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func RenderJSON(w http.ResponseWriter, obj interface{}) {
	json, _ := json.Marshal(obj)
	w.Write(json)
}

func RenderError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func Render404(w http.ResponseWriter) {
	http.Error(w, http.StatusText(404), http.StatusNotFound)
}

func strLen(input string) int {
	return len([]rune(input))
}

func makeQuery(begin string, getParams url.Values, queryParams map[string]string) (query string, params []interface{}) {
	query = begin + " "
	params = []interface{}{}

	hasWhere := false
	for key, value := range queryParams {
		param := getParams.Get(key)
		if strLen(param) > 0 {
			if !hasWhere {
				query += "where "
				hasWhere = true
			} else {
				query += "and "
			}
			if value == "like" {
				query += key + " " + value + " ? "
				param = "%" + param + "%"
			} else {
				query += key + " " + value + " " + "? "
			}
			params = append(params, param)
		}
	}

	start := getParams.Get("start")
	limit := getParams.Get("limit")

	query += " limit "
	if strLen(start) > 0 || strLen(limit) > 0 {

		if strLen(start) > 0 {
			query += "?, "
			params = append(params, start)
		}
		query += "?"
		if strLen(limit) > 0 {
			params = append(params, limit)
		} else {
			params = append(params, 100)
		}
	} else {
		query += "?"
		params = append(params, 100)
	}
	return
}

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
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
