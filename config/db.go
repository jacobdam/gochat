package config

import (
	"os"

	"github.com/jacobdam/gochat/datastore"
)

var DB = map[string]datastore.Config{
	"production":  datastore.Config{Url: os.Getenv("MONGOLAB_URI")},
	"development": datastore.Config{Url: "localhost", DB: "gochat"},
}
