package models

import (
	. "github.com/hiddenfounders/redimo/config"
	"github.com/globalsign/mgo"
)

type DataStore struct {
	Session *mgo.Session
}

func (ds *DataStore) GetCol(colName string) *mgo.Collection {
	return ds.Session.Copy().DB(DATABASE).C(colName)
}
