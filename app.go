package main

import (
	"fmt"
	"github.com/go-redis/redis"
	. "github.com/hiddenfounders/redimo/handlers"
	. "github.com/hiddenfounders/redimo/models"
	. "github.com/hiddenfounders/redimo/utils"
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
