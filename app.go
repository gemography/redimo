package main

import(
  "fmt"
  "github.com/globalsign/mgo"
  "github.com/go-redis/redis"
  "log"
)

type DataStore struct {
	session *mgo.Session
}

func (ds *DataStore) getCol(colName string) *mgo.Collection {
	return ds.session.Copy().DB("project").C("users")
}

func main() {
    fmt.Println("Starting...")
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
    const url = "mongodb://41.79.79.212:27017,41.79.79.212:27018,41.79.79.212:27019/?replicaSet=rs0"
    var ds DataStore
  	ds.session, err = mgo.Dial(url)
  	if err != nil {
  		log.Fatal(err)
  	} else {
      fmt.Println("MongoDB Connected :)")
    }
}
