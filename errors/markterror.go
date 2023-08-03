package errors

// MarktErr custom error type.
type MarktErr struct {
	Message    string                 `json:"message"`
	Status     int                    `json:"status"`
	Parameters map[string]interface{} `json:"parameters"`
}

// Error returns error message.
func (e MarktErr) Error() string {
	return e.Message
}

// Status returns the  HTTP status code.
func (e MarktErr) StatusCode() int {
	return e.Status
}

// Params returns the parameters of the RoomError (must include the error_code).
func (e MarktErr) Params() map[string]interface{} {
	return e.Parameters
}
