package controller

import (
	"encoding/json"
	"net/http"
	"tui/flightmarketplace/errors"
	"tui/flightmarketplace/model/dto"

	"github.com/gorilla/mux"
)

type FlightControllerInterface interface {
	// CreateFlight creates a new flight
	CreateFlight(w http.ResponseWriter, r *http.Request)
	// GetFlight returns a flight by id
	GetFlight(w http.ResponseWriter, r *http.Request)
	// ModifyFlight modifies an existing flight
	ModifyFlight(w http.ResponseWriter, r *http.Request)
	// DeleteFlight deletes an existing flight
	DeleteFlight(w http.ResponseWriter, r *http.Request)
	// GetAllFlights returns all flights
	GetAllFlights(w http.ResponseWriter, r *http.Request)
	// GetFlightCountsByDurationBuckets returns the number of flights per 60-minute buckets of flight durations
	GetFlightCountsByDurationBuckets(w http.ResponseWriter, r *http.Request)
}

// @Summary Create a new flight
// @Description Create a new flight based on the provided FlightDto
// @Tags flights
// @Accept json
// @Produce json
// @Param flightDto body dto.FlightDto true "FlightDto object"
// @Success 200 {object} model.Flight
// @Failure 400 {object} errors.MarktErr
// @Router /flights [post]
func (c *Controller) CreateFlight(w http.ResponseWriter, r *http.Request) {
	var flightDto dto.FlightDto
	if err := json.NewDecoder(r.Body).Decode(&flightDto); err != nil {
		errors.RespondWithError(w, http.StatusBadRequest, "Failed to parse request body")
		return
	}

	createdFlight := c.Ds.CreateFlight(flightDto)

	errors.RespondWithJSON(w, http.StatusOK, createdFlight)
}

// @Summary Get a flight by ID
// @Description Get a flight based on the provided flight ID
// @Tags flights
// @Produce json
// @Param id path string true "Flight ID"
// @Success 200 {object} model.Flight
// @Failure 400 {object} errors.MarktErr
// @Router /flights/{id} [get]
func (c *Controller) GetFlight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightID := params["id"]

	flight := c.Ds.GetFlightById(flightID)
	if flight == nil {
		errors.RespondWithError(w, http.StatusNotFound, "Flight not found")
		return
	}

	errors.RespondWithJSON(w, http.StatusOK, flight)
}

// @Summary Modify an existing flight
// @Description Modify an existing flight based on the provided FlightDto and flight ID
// @Tags flights
// @Accept json
// @Produce json
// @Param id path string true "Flight ID"
// @Param flightDto body dto.FlightDto true "FlightDto object"
// @Success 200 {object} string
// @Failure 400 {object} errors.MarktErr
// @Router /flights/{id} [put]
func (c *Controller) ModifyFlight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightID := params["id"]

	var flightDto dto.FlightDto
	if err := json.NewDecoder(r.Body).Decode(&flightDto); err != nil {
		errors.RespondWithError(w, http.StatusBadRequest, "Failed to parse request body")
		return
	}

	c.Ds.ModifyFlight(flightID, flightDto)

	errors.RespondWithJSON(w, http.StatusOK, "Flight modified successfully")
}

// @Summary Delete a flight by ID
// @Description Delete a flight based on the provided flight ID
// @Tags flights
// @Produce json
// @Param id path string true "Flight ID"
// @Success 200 {object} string
// @Failure 400 {object} errors.MarktErr
// @Router /flights/{id} [delete]
func (c *Controller) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightID := params["id"]

	c.Ds.DeleteFlight(flightID)

	errors.RespondWithJSON(w, http.StatusOK, "Flight deleted successfully")
}

// @Summary Get all flights
// @Description Get a list of all flights
// @Tags flights
// @Produce json
// @Success 200 {array} model.Flight
// @Router /flights [get]
func (c *Controller) GetAllFlights(w http.ResponseWriter, r *http.Request) {
	flights := c.Ds.GetAllFlights()

	errors.RespondWithJSON(w, http.StatusOK, flights)
}

// @Summary Get flight counts by duration buckets
// @Description Get the number of flights per 60-minute buckets of flight durations
// @Tags flights
// @Produce json
// @Success 200 {object} map[string]int
// @Router /flight-counts [get]
func (c *Controller) GetFlightCountsByDurationBuckets(w http.ResponseWriter, r *http.Request) {
	flightCounts := c.Ds.GetFlightCountsByDurationBuckets()

	errors.RespondWithJSON(w, http.StatusOK, flightCounts)
}
