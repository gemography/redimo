package main

import (
	. "./handlers"
	. "./models"
	. "./utils"
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	RoutineErrors := make(chan string)
	var ds DataStore
	ConnectRedis()
	ds = ConnectMongo()
	go HandleUsers(ds.GetCol("users"), RoutineErrors)
}
