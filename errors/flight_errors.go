package errors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ErrUsernameNotExist() *MarktErr {

	code := "ERR_USERNAME_NOT_EXIST"
	msg := fmt.Sprintf("%s: tried to make a request using a non existing username.", code)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

func ErrNotFound() *MarktErr {

	code := "ERR_NO_FLIGHT_RECORD_FOUND_IN_DB"
	msg := fmt.Sprintf("%s: tried to make a request using a non existing flight.", code)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

// Controller errors
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		// If there's an error while marshaling the data, respond with an internal server error
		RespondWithError(w, http.StatusInternalServerError, "Failed to marshal response data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		// If there's an error while writing the response, log the error
		log.Println("Failed to write response:", err)
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	errorResponse := map[string]interface{}{
		"error": errorMessage,
	}

	response, err := json.Marshal(errorResponse)
	if err != nil {
		// If there's an error while marshaling the error response, write a simple error message to the response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		// If there's an error while writing the response, log the error
		log.Println("Failed to write response:", err)
	}
}
