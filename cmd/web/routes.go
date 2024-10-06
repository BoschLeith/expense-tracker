package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.home)
	return mux
}
