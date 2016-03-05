package main

import (
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/jacobdam/gochat/config"
	"github.com/jacobdam/gochat/datastore"
)

var APP_ENV string

func init() {
	APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV == "" {
		APP_ENV = "development"
	}
}

func createDS() *datastore.DataStore {
	ds, _ := datastore.New(config.DB[APP_ENV])
	return ds
}

func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/products", ProductsHandler)

	return r
}

func injectDsHandler(h http.Handler, ds *datastore.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dsCopy := ds.Copy()
		defer dsCopy.Close()
		context.Set(r, "ds", dsCopy)
		h.ServeHTTP(w, r)
	})
}

func CreateApp() http.Handler {
	ds := createDS()
	r := createRouter()
	var h http.Handler = r
	h = injectDsHandler(h, ds)
	h = handlers.RecoveryHandler()(h)
	h = handlers.LoggingHandler(os.Stdout, h)

	return h
}
