package handler

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	html, err := Asset("assets/html/index.html")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(html)
}
