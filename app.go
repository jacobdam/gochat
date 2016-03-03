package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateApp() http.Handler {
	r := mux.NewRouter()

	ConfigRoutes(r)

	return r
}
