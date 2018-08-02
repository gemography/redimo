package models

import (
	"github.com/globalsign/mgo"
	. "github.com/hiddenfounders/redimo/config"
)

type DataStore struct {
	Session *mgo.Session
}

func (ds *DataStore) GetCol(colName string) *mgo.Collection {
	return ds.Session.Copy().DB(DATABASE).C(colName)
}
