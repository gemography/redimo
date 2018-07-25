package main

import(
  "fmt"
  "github.com/globalsign/mgo"
  "github.com/go-redis/redis"
  "log"
  "strings"
)

type DataStore struct {
	session *mgo.Session
}

func (ds *DataStore) getCol(colName string) *mgo.Collection {
	return ds.session.Copy().DB("project").C("users")
}

const (
  PROTOCOL = "mongodb://"
	REPLICASET = "/?replicaSet="
)

func createURI() string {
  hosts := []string{"41.79.79.212:27017", "41.79.79.212:27018", "41.79.79.212:27019"};
  var replicaSetName = "rs0"
  var URI = PROTOCOL + strings.Join(hosts, ",") + REPLICASET + replicaSetName
  return URI
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
    var ds DataStore
  	ds.session, err = mgo.Dial(createURI())
  	if err != nil {
  		log.Fatal(err)
  	} else {
      fmt.Println("MongoDB Connected :)")
    }
}
