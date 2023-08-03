package dataservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	err_handler "tui/flightmarketplace/errors"
	"tui/flightmarketplace/model"
	"tui/flightmarketplace/model/dto"

	"github.com/gofrs/uuid"
)

type FlightDataservice interface {
	CreateFlight(flightDto dto.FlightDto) model.Flight
	ModifyFlight(flightID string, flightDto dto.FlightDto) uuid.UUID
	DeleteFlight(flightID string) uuid.UUID
	GetFlightById(flightID string) *model.Flight
	GetAllFlights() []model.Flight
	GetFlightCountsByDurationBuckets() map[string]int
}

// CreateFlight converts the FlightDto to a model.Flight and creates a new flight record in the database
func (ds *DataGatewayService) CreateFlight(flightDto dto.FlightDto) model.Flight {
	// Validation code ...

	// Convert the FlightDto to model.Flight
	departureTime, err := time.Parse(time.RFC3339, flightDto.DepartureDateTime)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid departure date time"))
	}

	duration, err := time.ParseDuration(flightDto.Duration)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid duration"))
	}

	price, err := strconv.ParseFloat(flightDto.Price, 64)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid price"))
	}

	availableSeats, err := strconv.Atoi(flightDto.AvailableSeats)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid available seats"))
	}

	flight := model.Flight{
		Origin:            flightDto.Origin,
		Destination:       flightDto.Destination,
		DepartureDateTime: departureTime,
		ArrivalDateTime:   departureTime.Add(duration),
		Price:             price,
		AvailableSeats:    availableSeats,
		// You can set other fields of the model.Flight as needed
	}

	// Now, create the flight in the database using your flight database implementation
	createdFlight, err := ds.Db.CreateFlight(&flight)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadSyntax())
	}

	return *createdFlight
}

func (ds *DataGatewayService) ModifyFlight(flightID string, flightDto dto.FlightDto) uuid.UUID {
	// Validation code ...

	// Convert the FlightDto to model.Flight
	departureTime, err := time.Parse(time.RFC3339, flightDto.DepartureDateTime)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrServerError(http.StatusText(http.StatusInternalServerError)))
		return uuid.Nil
	}

	duration, err := time.ParseDuration(flightDto.Duration)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrServerError(http.StatusText(http.StatusInternalServerError)))
		return uuid.Nil
	}

	price, err := strconv.ParseFloat(flightDto.Price, 64)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrServerError(http.StatusText(http.StatusInternalServerError)))
		return uuid.Nil
	}

	availableSeats, err := strconv.Atoi(flightDto.AvailableSeats)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrServerError(http.StatusText(http.StatusInternalServerError)))
		return uuid.Nil
	}
	uuidd, err := uuid.FromString(flightID)
	if err != nil {
		// Handle error if the UUID string is not valid
		fmt.Println("Error parsing UUID:", err)
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid flight ID"))
		return uuid.Nil
	}
	flight := model.Flight{
		ID:                uuidd, // Convert flightID to uuid.UUID
		Origin:            flightDto.Origin,
		Destination:       flightDto.Destination,
		DepartureDateTime: departureTime,
		ArrivalDateTime:   departureTime.Add(duration),
		Price:             price,
		AvailableSeats:    availableSeats,
		// You can set other fields of the model.Flight as needed
	}

	// Now, modify the flight in the database using your flight database implementation
	id, err := ds.Db.ModifyFlight(&flight)
	if err != nil {
		newErr := err_handler.ErrServerError(err.Error())
		err_handler.ThrowError(newErr)
	}
	return id
}

func (ds *DataGatewayService) DeleteFlight(flightID string) uuid.UUID {
	// Validation code ...
	uuidd, err := uuid.FromString(flightID)
	if err != nil {
		// Handle error if the UUID string is not valid
		fmt.Println("Error parsing UUID:", err)
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid flight ID"))
		return uuid.Nil
	}
	// Now, delete the flight from the database using your flight database implementation
	id, err := ds.Db.DeleteFlight(uuidd)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid flight ID"))
		return uuid.Nil
	}
	return id
}

func (ds *DataGatewayService) GetFlightById(flightID string) *model.Flight {
	// Validation code ...
	uuid, err := uuid.FromString(flightID)
	if err != nil {
		// Handle error if the UUID string is not valid
		fmt.Println("Error parsing UUID:", err)
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid flight ID"))
		return nil
	}
	// Get the flight from the database using flightID
	flight, err := ds.Db.GetFlightById(uuid)
	if err != nil {
		err_handler.ThrowError(err_handler.ErrBadRequest("Invalid flight ID"))
		return nil
	}

	return flight
}

func (ds *DataGatewayService) GetAllFlights() []model.Flight {
	// Get all flights from the database using your flight database implementation
	flights, err := ds.Db.GetAllFlights()
	if err != nil {
		err_handler.ThrowError(err)
	}

	return flights
}

func (ds *DataGatewayService) GetFlightCountsByDurationBuckets() map[string]int {
	// Call the GetFlightCountsByDurationBuckets function in the database
	flightCounts, err := ds.Db.GetFlightCountsByDurationBuckets()
	if err != nil {
		err_handler.ThrowError(err)
	}

	return flightCounts
}
