package models

import (
	"github.com/globalsign/mgo/bson"
)

type ChangeDocument struct {
	ResumeToken struct {
		Data bson.ObjectId `bson:"_data"`
	} `bson:"_id"`
	OperationType string   `bson:"operationType"`
	FullDocument  bson.Raw `bson:"fullDocument"`
	DocumentKey   struct {
		ID bson.ObjectId `bson:"_id"`
	} `bson:"documentKey"`
}
