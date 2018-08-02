package handlers

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/go-redis/redis"
	. "github.com/hiddenfounders/redimo/models"
	. "github.com/hiddenfounders/redimo/utils"
	"log"
)

func HandleUsers(c *mgo.Collection, RoutineErrors chan string, client *redis.Client) {
	token := c.Name + "ResumeToken"
	var changeStream = ListenOnPipeline(c, RoutineErrors, client, token)
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
