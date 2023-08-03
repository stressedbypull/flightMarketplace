package controller

import "tui/flightmarketplace/dataservice"

type ControllerInterface interface {
	FlightControllerInterface
	//UserControllerInterface
	//OrderControllerInterface
}

type Controller struct {
	Ds dataservice.DataServiceInterface
	//Mutex *sync.Mutex
}
