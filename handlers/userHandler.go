package handlers

import (
	. "../models"
	. "../utils"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/go-redis/redis"
	"log"
)

func HandleUsers(c *mgo.Collection, RoutineErrors chan string) {
	var client *redis.Client
	var changeStream = ListenOnPipeline(c, RoutineErrors, client)
	token := c.Name + "ResumeToken"
	changeDoc := ChangeDocument{}
	User := User{}
	var x interface{}
	var err error
	for {
		for changeStream.Next(&changeDoc) {
			changeDoc.FullDocument.Unmarshal(&User)
			x, err = bson.Marshal(changeStream.ResumeToken())
			fmt.Printf("%+v \n", User)
			if err != nil {
				log.Fatal(err)
			}
			client.Set(token, x, 0)
		}
	}
}
