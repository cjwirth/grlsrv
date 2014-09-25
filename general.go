package main

import (
	"encoding/json"
	"net/http"
	"net/url"
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
