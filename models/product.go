package models

// import "gopkg.in/mgo.v2/bson"

type Product struct {
	// Id    bson.ObjectId `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
