package errors

import (
	"fmt"
	"net/http"
)

// ErrBadSyntax is raised when user provides a form or body with missing or invalid fields.
func ErrBadSyntax() *MarktErr {

	code := "ERR_BAD_SYNTAX"
	msg := fmt.Sprintf("%s: tried to make a request using body or form with missing or invalid fields.", code)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

// ErrUnauthorized is raised when user is not authorized to use the API.
func ErrUnauthorized() *MarktErr {

	code := "USER_NOT_AUTHORIZED"
	msg := fmt.Sprintf("%s: tried to make a request without the proper authorization.", code)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusUnauthorized,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

// ErrMalformedBody is raised when the call's body cannot be decoded.
func ErrMalformedBody() *MarktErr {

	code := "ERR_MALFORMED_BODY"
	msg := fmt.Sprintf("%s: tried to make a request using a json body that could not be decoded.", code)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusUnprocessableEntity,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

// ErrServerError is raised when server breaks for internal reasons.
func ErrServerError(message string) *MarktErr {

	code := "ERR_INTERNAL_SERVER_ERROR"
	msg := fmt.Sprintf("%s: an error was thrown during internal server processes. Error message: '%s'.", code, message)

	return &MarktErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

// ErrServerError is raised when server breaks for internal reasons.
func ErrBadRequest(message string) *MarktErr {

	code := "ERR_BAD_REQUEST"
	msg := message
	if msg == "" {
		msg = fmt.Sprintf("%s: an error was thrown during bad request. Error message: '%s'.", code, message)
	}
	return &MarktErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
