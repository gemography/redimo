package utils

import (
	. "../config"
	. "../models"
	"fmt"
	"github.com/globalsign/mgo"
	"log"
	"strings"
)

// CreateURI() is a function that concatenates multiple hosts
// to create a MongoDB valid URI for replica sets.
func CreateURI() string {
	hosts := []string{"41.79.79.212:27017", "41.79.79.212:27018", "41.79.79.212:27019"}
	var replicaSetName = "rs0"
	var URI = PROTOCOL + strings.Join(hosts, ",") + REPLICASET + replicaSetName
	return URI
}

// ConnectMongo() is a function that connects to Mongo and prints
// a success message and it returns a a valid DataStore
// or an error message in case of failure.
func ConnectMongo() DataStore {
	var ds DataStore
	var err error
	ds.Session, err = mgo.Dial(CreateURI())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("MongoDB Connected :)")
	}
	return ds
}
