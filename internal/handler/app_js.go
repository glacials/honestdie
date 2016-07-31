package handler

import (
	"net/http"
)

func AppJS(w http.ResponseWriter, r *http.Request) {
	html, err := Asset("assets/js/app.js")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(html)
}
