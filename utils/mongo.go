package utils

import (
	. "../config"
	. "../models"
	"fmt"
	"github.com/globalsign/mgo"
	"log"
	"strings"
)

func CreateURI() string {
	hosts := []string{"41.79.79.212:27017", "41.79.79.212:27018", "41.79.79.212:27019"}
	var replicaSetName = "rs0"
	var URI = PROTOCOL + strings.Join(hosts, ",") + REPLICASET + replicaSetName
	return URI
}

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
