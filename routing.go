package main

import "github.com/gorilla/mux"

func ConfigRoutes(r *mux.Router) {
	r.Handle("/products", ProductsHandler)
}
