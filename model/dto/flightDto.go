package dto

// FlightDto is a data transfer object for flights
// Frontend is supposed to use this to create new flights
type FlightDto struct {
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	DepartureDateTime string `json:"departureDateTime"`
	Duration          string `json:"duration"`
	Price             string `json:"price"`
	AvailableSeats    string `json:"availableSeats"`
}
