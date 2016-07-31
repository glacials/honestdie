package handler

import (
	"net/http"
)

func BodyCSS(w http.ResponseWriter, r *http.Request) {
	html, err := Asset("assets/css/body.css")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(html)
}
