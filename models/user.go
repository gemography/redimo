package models

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id"  json:"id"`
	Username string        `bson:"username" json:"username"`
	Email    string        `bson:"email" json:"email"`
}
