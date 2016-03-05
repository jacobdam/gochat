package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"

	"github.com/jacobdam/gochat/datastore"
	"github.com/jacobdam/gochat/models"
)

var ProductsHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	var ds *datastore.DataStore = context.Get(r, "ds").(*datastore.DataStore)
	var products []models.Product
	ds.QueryAllProducts(&products)
	json.NewEncoder(w).Encode(products)
}
