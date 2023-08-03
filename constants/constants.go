package constants

// Different color codes useful for printing logs
const (
	Green    = "\033[0;32m"
	Red      = "\033[0;31m"
	ColorOff = "\033[0m"
)

const (
	//HTTP RESPONSE Messages
	GENERIC_ERROR_CODE = "GENERIC"
	NOT_FOUND_CODE     = "NOT_FOUND"
	BAD_REQUEST_CODE   = "BAD_REQUEST"

	// API Err Messages
	GENERIC_ERROR_MSG        = "Something was gone wrong!"
	INVALID_ID_FLIGHT        = "Invalid 'id' in request"
	INVALID_FLIGHT_DURATION  = "Invalid 'duration' in request"
	INVALID_BODY_REQUEST_MSG = "Invalid body request"

	//Auth constant err messages
	USER_NOT_AUTHORIZED         = "USER_NOT_AUTHORIZED"
	CLAIM_NOT_RECOGNIZED        = "CLAIM_NOT_RECOGNIZED"
	ERR_USER_NOT_FOUND          = "ERR_USER_NOT_FOUND"
	ERR_GROUP_NOT_FOUND         = "ERR_GROUP_NOT_FOUND"
	ERR_CONFIRM_TOKEN_NOT_FOUND = "ERR_CONFIRM_TOKEN_NOT_FOUND"
	ERR_EXPIRED_CONFIRM_TOKEN   = "ERR_EXPIRED_CONFIRM_TOKEN"

	//Groups for users authentication and verification
	AdminGroupManage    = "flight-admin"
	MarketerGroupManage = "flight-marketer"
	CustomerGroupManage = "flight-customer"
)
