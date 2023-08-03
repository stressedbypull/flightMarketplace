package router

import (
	"tui/flightmarketplace/controller"

	"github.com/gorilla/mux"
)

func setupFlightRouter(router *mux.Router, c controller.ControllerInterface) {
	// Register the routes and handlers
	//router.HandleFunc("/flights", marketerAuthMiddleware(c.CreateFlight)).Methods("POST")
	router.HandleFunc("/flights", c.CreateFlight).Methods("POST")
	router.HandleFunc("/flights/{id}", c.GetFlight).Methods("GET")
	router.HandleFunc("/flights/{id}", c.ModifyFlight).Methods("PUT")
	//router.HandleFunc("/flights/{id}", marketerAuthMiddleware(c.ModifyFlight)).Methods("PUT")
	router.HandleFunc("/flights/{id}", c.DeleteFlight).Methods("DELETE")
	router.HandleFunc("/flights", c.GetAllFlights).Methods("GET")
	router.HandleFunc("/flight-counts", c.GetFlightCountsByDurationBuckets).Methods("GET")
}
