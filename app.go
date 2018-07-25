package main

import(
  "fmt"
  "github.com/globalsign/mgo"
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
    const url = "mongodb://41.79.79.212:27017,41.79.79.212:27018,41.79.79.212:27019/?replicaSet=rs0"
    var ds DataStore
  	ds.session, err = mgo.Dial(url)
  	if err != nil {
  		log.Fatal(err)
  	} else {
      fmt.Println("Database Connected :)")
    }
}
