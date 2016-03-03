package main

import "net/http"
import "encoding/json"
import "github.com/jacobdam/gochat/models"

var ProductsHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.AllProducts)
}
