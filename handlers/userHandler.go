package handlers

import (
	. "../models"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client

func HandleUsers(c *mgo.Collection, RoutineErrors chan string) {
	pipeline := []bson.M{}
	ResumeToken := bson.Raw{}
	options := mgo.ChangeStreamOptions{}
	val, err := client.Get("userResumeToken").Bytes()
	if err == redis.Nil {
		fmt.Println("no previous resume token")
	} else if err != nil {
		log.Fatal(err)
	} else {
		err := bson.Unmarshal(val, &ResumeToken)
		if err != nil {
			log.Fatal(err)
		}
		options.ResumeAfter = &ResumeToken
	}

	changeStream, err := c.Watch(pipeline, options)
	if err != nil {
		if err := changeStream.Close(); err != nil {
			log.Fatal(err)
			RoutineErrors <- "error in user handler"
		}
	}
	changeDoc := ChangeDocument{}

	User := User{}
	var x interface{}
	var e error
	for {
		for changeStream.Next(&changeDoc) {
			changeDoc.FullDocument.Unmarshal(&User)
			x, e = bson.Marshal(changeStream.ResumeToken())
			fmt.Printf("%+v \n", User)
			if e != nil {
				log.Fatal(e)
			}
			client.Set("userResumeToken", x, 0)
		}
	}
}
