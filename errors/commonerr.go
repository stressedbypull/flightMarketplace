package errors

import "net/http"

// ErrServerError is raised when server breaks for internal reasons.
var ErrServerError = &MarktErr{message: "ERR_INTERNAL_SERVER_ERROR", status: http.StatusInternalServerError}

var ErrUnauthorized = &MarktErr{message: "User not authorized", status: http.StatusUnauthorized}

// ErrMissingToken is raised when request does not contain a jwt for an API which requires authentication.
var ErrMissingToken = &MarktErr{message: "ERR_MISSING_TOKEN", status: http.StatusUnauthorized}

var ErrInvalidToken = &MarktErr{message: "ERR_INVALID_TOKEN", status: http.StatusUnauthorized}

var ErrInContextRequest = &MarktErr{message: "ERR_IN_CONTEXT_REQUEST", status: http.StatusInternalServerError}

// ErrBadSyntax is raised when user provides a form or body with missing or invalid fields.
var ErrBadSyntax = &MarktErr{message: "ERR_BAD_SYNTAX", status: http.StatusBadRequest}

var ErrNotFound = &MarktErr{message: "ERR_NO_RECORD_FOUND_IN_DB", status: http.StatusNotFound}

var ErrInvalidParamErr = &MarktErr{message: "ERR_INVALID_PARAM", status: http.StatusBadRequest}

// ErrOperationForbidden raised when application logic blocks an operation (no 401) for a user, for example because
// some precondition is not satisfied.
var ErrOperationForbidden = &MarktErr{message: "ERR_OPERATION_FORBIDDEN", status: http.StatusForbidden}

var ErrUsernameNotExist = &MarktErr{message: "ERR_USERNAME_NOT_EXIST", status: http.StatusBadRequest}
