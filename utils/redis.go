package utils

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/go-redis/redis"
	"log"
)

// ConnectRedis() is a function that checks if
// Redis server is on, and it prints success image in
// case of success and an error message is case of
// failure.
func ConnectRedis() *redis.Client {
	var err error
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Redis Connected :)")
	}
	return client
}

// ListenOnPipeline() is a function that takes as arguments Mongo Collection,
//RoutineErrors and redis client, and it returns Mongo Change Stream. It listens
// on the pipeline for any changes and it checks for ResumeToken to start from
// it in case of previous failure.
func ListenOnPipeline(c *mgo.Collection, RoutineErrors chan string, client *redis.Client, token string) *mgo.ChangeStream {
	pipeline := []bson.M{}
	ResumeToken := bson.Raw{}
	options := mgo.ChangeStreamOptions{}
	val, err := client.Get(token).Bytes()
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
			RoutineErrors <- "error in " + c.Name + " handler"
		}
	}
	return changeStream
}
