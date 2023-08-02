package main

import (
	"fmt"
	"log"
	"tui/flightmarketplace/config"
	"tui/flightmarketplace/database"
)

func main() {
	//configure settings from json
	config.SetupConfig()

	//connect to database
	postgres, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//create tables if not present
	err = postgres.ConfigureTables()
	if err != nil {
		log.Fatal(err)
	}

	//connect to a redis instance
	//redisClient, err := redis.Connect()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//create a new router
	//router := router.NewRouter()
	//Listen and serve
	//log.Fatal(http.ListenAndServe(":"+config.Config.Server.Port, router))
	fmt.Printf("Hola mundo")
}
