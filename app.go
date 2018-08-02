package main

import (
	. "github.com/hiddenfounders/redimo/handlers"
	. "github.com/hiddenfounders/redimo/models"
	. "github.com/hiddenfounders/redimo/utils"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	fmt.Println("Starting...")
	RoutineErrors := make(chan string)
	var ds DataStore
	client = ConnectRedis()
	ds = ConnectMongo()
	go HandleUsers(ds.GetCol("users"), RoutineErrors, client)
	<-RoutineErrors
}
