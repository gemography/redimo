package main

import (
	. "./utils"
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	ConnectRedis()
	ConnectMongo()
}
