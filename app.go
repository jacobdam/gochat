package main

import (
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/jacobdam/gochat/datastore"
)

func createDS() *datastore.DataStore {
	ds, _ := datastore.New(datastore.Config{"localhost:27017", "gochat"})
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
