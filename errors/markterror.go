package errors

import (
	"log"
	"tui/flightmarketplace/constants"
)

// MarktErr custom error type.
type MarktErr struct {
	message string
	status  int // HTTP status codes as registered with IANA.
}

// Error returns error message.
func (e MarktErr) Error() string {
	return e.message
}

// Status returns HTTP status code as registered with IANA.
func (e MarktErr) Status() int {
	return e.status
}

// NewError returns an initialized MarktErr.
func NewError(message string, status int) error {
	return &MarktErr{
		message: message,
		status:  status,
	}
}

func (e *MarktErr) setMessage(message string) {
	e.message = message
}

// CheckError checks if the error is not nil, and eventually panics.
func CheckError(err error, msg string) {
	if err != nil {
		log.Printf("%s:\n%s", msg, err.Error())
		panic(ErrServerError)
	}
}

func InvalidParamErr(paramName string) *MarktErr {
	err := ErrInvalidParamErr
	err.setMessage(err.message + " " + paramName)
	return err
}

// ThrowError panics an application error that will be recovered in middleware.
func ThrowError(err error) {
	panic(err)
}

// LogError displays an error message with the right format.
func LogError(message string) {
	log.Println(constants.Red + "ERROR: " + message + constants.ColorOff)
}
