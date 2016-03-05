package datastore

import "gopkg.in/mgo.v2"

type Config struct {
	Url string
	DB  string
}

type DataStore struct {
	session *mgo.Session
	config  Config
}

type ConfigError struct {
	error
}

func New(cf Config) (ds *DataStore, err error) {
	session, e := mgo.Dial(cf.Url)
	if e != nil {
		err = ConfigError{e}
		return
	}

	ds = &DataStore{session, cf}
	return
}

func (ds *DataStore) Copy() *DataStore {
	return &DataStore{ds.session.Copy(), ds.config}
}

func (ds *DataStore) Close() {
	ds.session.Close()
}

func (ds *DataStore) db() *mgo.Database {
	return ds.session.DB(ds.config.DB)
}
