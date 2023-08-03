package main

import (
	"log"
	"net/http"
	"tui/flightmarketplace/config"
	"tui/flightmarketplace/controller"
	"tui/flightmarketplace/database"
	"tui/flightmarketplace/dataservice"
	"tui/flightmarketplace/router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	routerMux := mux.NewRouter()
	// Create the dataservice
	ds := &dataservice.DataGatewayService{
		Db: postgres,
	}
	// Create the controller
	c := &controller.Controller{Ds: ds}
	router.SetupRouter(routerMux, c)

	println("Start flight-marketplace: server on port ", config.Config.Server.Port)

	//Listen and serve
	log.Fatal(http.ListenAndServe(":"+config.Config.Server.Port, cors.AllowAll().Handler(routerMux)))
}
