package datastore

import "gopkg.in/mgo.v2/bson"
import "github.com/jacobdam/gochat/models"

func (ds *DataStore) QueryAllProducts(result *[]models.Product) error {
	return ds.db().C("products").Find(bson.M{}).All(result)
}
