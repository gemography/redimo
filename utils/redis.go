package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func ConnectRedis() {
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
}
