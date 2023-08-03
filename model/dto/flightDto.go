package dto

// FlightDto is a data transfer object for flights
// Frontend is supposed to use this to create new flights
type FlightDto struct {
	Origin            string `json:"origin" example:"New York"`
	Destination       string `json:"destination" example:"Los Angeles"`
	DepartureDateTime string `json:"departureDateTime" example:"2023-08-03T12:00:00Z"`
	Duration          string `json:"duration" example:"3h30m0s"`
	Price             string `json:"price" example:"250.50"`
	AvailableSeats    string `json:"availableSeats" example:"150"`
}
