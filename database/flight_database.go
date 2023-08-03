package database

import (
	"fmt"
	"log"
	err_handler "tui/flightmarketplace/errors"
	"tui/flightmarketplace/model"

	"github.com/gofrs/uuid"
)

type FlightDatabaseInterface interface {
	//CRUD OPERATION
	// CreateFlight creates a new flight
	CreateFlight(flight *model.Flight) (*model.Flight, *err_handler.MarktErr)
	// ModifyFlight modifies an existing flight
	ModifyFlight(flight *model.Flight) (uuid.UUID, *err_handler.MarktErr)
	// DeleteFlight deletes an existing flight
	DeleteFlight(flightID uuid.UUID) (uuid.UUID, *err_handler.MarktErr)
	// GetFlights returns a flight by its ID
	GetFlightById(flightID uuid.UUID) (*model.Flight, *err_handler.MarktErr)
	//Get all flights
	GetAllFlights() ([]model.Flight, *err_handler.MarktErr)

	// QUERY OPERATION FOR SEARCHING FLIGHTS
	// GetFlightCountsByDurationBuckets returns the number of flights per 60-minute buckets of flight durations
	GetFlightCountsByDurationBuckets() (map[string]int, *err_handler.MarktErr)
}

func (pg PostgresClient) CreateFlight(flight *model.Flight) (*model.Flight, *err_handler.MarktErr) {
	result := pg.DB.Create(&flight)
	if result.Error != nil {
		log.Printf("Error while saving Flight Object, caused by %v", result.Error.Error())
		return nil, err_handler.ErrBadSyntax() // Return the custom error using the ErrBadSyntax() function
	}
	log.Printf("Flight saved with id %v", flight.ID)
	return flight, nil
}

// ModifyFlight modifies an existing flight
func (pg PostgresClient) ModifyFlight(flight *model.Flight) (uuid.UUID, *err_handler.MarktErr) {
	result := pg.DB.Save(&flight)
	if result.Error != nil {
		log.Printf("Error while modifying Flight Object, caused by %v", result.Error.Error())
		return uuid.Nil, err_handler.ErrServerError(result.Error.Error())
	}
	log.Printf("Flight with ID %v modified successfully", flight.ID)
	return flight.ID, nil
}

// DeleteFlight deletes an existing flight
func (pg PostgresClient) DeleteFlight(flightID uuid.UUID) (uuid.UUID, *err_handler.MarktErr) {
	// Retrieve the flight by its ID
	var flight model.Flight
	if err := pg.DB.First(&flight, flightID).Error; err != nil {
		log.Printf("Error while fetching Flight with ID %v, caused by %v", flightID, err.Error())
		return uuid.Nil, err_handler.ErrNotFound()
	}

	// Delete the retrieved flight
	result := pg.DB.Delete(&flight)
	if result.Error != nil {
		log.Printf("Error while deleting Flight Object, caused by %v", result.Error.Error())
		return uuid.Nil, err_handler.ErrServerError(result.Error.Error())
	}

	log.Printf("Flight with ID %v deleted successfully", flightID)
	return flightID, nil
}

// GetFlightById returns a flight by its ID
func (pg PostgresClient) GetFlightById(flightID uuid.UUID) (*model.Flight, *err_handler.MarktErr) {
	var flight model.Flight
	result := pg.DB.First(&flight, flightID)
	if result.Error != nil {
		log.Printf("Error while fetching Flight with ID %v, caused by %v", flightID, result.Error.Error())
		return nil, err_handler.ErrNotFound()
	}
	return &flight, nil
}

// GetAllFlights gets all flights
func (pg PostgresClient) GetAllFlights() ([]model.Flight, *err_handler.MarktErr) {
	var flights []model.Flight
	result := pg.DB.Find(&flights)
	if result.Error != nil {
		log.Printf("Error while fetching all flights, caused by %v", result.Error.Error())
		return nil, err_handler.ErrServerError(result.Error.Error())
	}
	return flights, nil
}

// GetFlightCountsByDurationBuckets returns the number of flights per 60-minute buckets of flight durations
func (pg PostgresClient) GetFlightCountsByDurationBuckets() (map[string]int, *err_handler.MarktErr) {
	bucketSize := 60 // Duration bucket size in minutes

	type BucketResult struct {
		Bucket string
		Count  int
	}

	var results []BucketResult
	queryCategorizedFlights := fmt.Sprintf(
		"CASE WHEN (EXTRACT(EPOCH FROM (arrival_date_time - departure_date_time)) / 60) < %d"+
			" THEN '0-%d mins' ELSE (FLOOR(EXTRACT(EPOCH FROM (arrival_date_time - departure_date_time)) / 60) / %d * %d) || '-' ||"+
			" (FLOOR(EXTRACT(EPOCH FROM (arrival_date_time - departure_date_time)) / 60) / %d * %d + %d)"+
			" || ' mins' END as duration_bucket, COUNT(*) as count", bucketSize, bucketSize, bucketSize, bucketSize, bucketSize, bucketSize, bucketSize)
	// GORM query to group flights by duration buckets and count the number of flights in each bucket
	result := pg.DB.
		Model(&model.Flight{}).
		Select(queryCategorizedFlights).
		Group("duration_bucket").
		Find(&results)

	if result.Error != nil {
		log.Printf("Error while fetching flight counts by duration buckets, caused by %v", result.Error.Error())
		return nil, err_handler.ErrServerError(result.Error.Error())
	}

	// Convert the result to a map
	flightCounts := make(map[string]int)
	for _, r := range results {
		flightCounts[r.Bucket] = r.Count
	}

	return flightCounts, nil
}
