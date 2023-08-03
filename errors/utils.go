package errors

import (
	"log"
	"time"
	"tui/flightmarketplace/constants"
)

// CheckDbError check if the error is not nil, and eventually panics.
func CheckDbError(err error) {
	if err != nil {
		log.Printf("DB error: %s", err.Error())
		ThrowError(ErrServerError(err.Error()))
	}
}

// CheckRowError check if the error is not nil and different from r.ErrEmptyResult, and eventually panics.
func CheckRowError(err error) {
	if err != nil {
		log.Printf("DB error: %s", err)
		ThrowError(ErrServerError(err.Error()))
	}
}

// ThrowError panics an application error that will be recovered in middleware.
func ThrowError(err *MarktErr) {
	panic(err)
}

// FormatDateForLogging formats a date in timestamp form, turning into a date string suitable for logging.
func FormatTsForLogging(ts int64) string {
	return time.Unix(ts/1000, 0).Format(time.RFC822Z)
}

// NewError returns an initialized MarktErr.
func NewError(message string, status int) error {
	return &MarktErr{
		Message: message,
		Status:  status,
	}
}

// CheckError checks if the error is not nil, and eventually panics.
func CheckError(err error, msg string) {
	if err != nil {
		log.Printf("%s:\n%s", msg, err.Error())
		panic(ErrServerError)
	}
}

// LogError displays an error message with the right format.
func LogError(message string) {
	log.Println(constants.Red + "ERROR: " + message + constants.ColorOff)
}
