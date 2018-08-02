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
func CreateURI(withPassword bool) string {
	hosts := []string{"localhost:28017", "localhost:28018", "localhost:28019"}
	var replicaSetName = "rs0"
	var URI = ""
	if withPassword {
		URI = PROTOCOL + MGUSER + ":" + MGPASS + "@" + strings.Join(hosts, ",") + REPLICASET + replicaSetName + AUTHSOURCE + DATABASE
	} else {
		URI = PROTOCOL + strings.Join(hosts, ",") + REPLICASET + replicaSetName
	}
	return URI
}

// ConnectMongo() is a function that connects to Mongo and prints
// a success message and it returns a a valid DataStore
// or an error message in case of failure.
func ConnectMongo() DataStore {
	var ds DataStore
	var err error
	bool withPassword = false
	ds.Session, err = mgo.Dial(CreateURI(withPassword))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("MongoDB Connected :)")
	}
	return ds
}
