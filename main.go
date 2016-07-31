package main

import (
	"fmt"
	"net/http"

	"github.com/glacials/honestdie/handler"
	"golang.org/x/net/websocket"
)

func main() {
	mux := buildMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("can't start server: %s", err))
	}
}

func buildMux() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handler.Root))
	mux.Handle("/assets/css/body.css", http.HandlerFunc(handler.BodyCSS))
	mux.Handle("/assets/js/app.js", http.HandlerFunc(handler.AppJS))
	mux.Handle("/websocket", websocket.Handler(handler.Websocket))

	return mux
}
