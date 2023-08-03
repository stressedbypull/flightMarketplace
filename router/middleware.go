package router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"tui/flightmarketplace/config"
	"tui/flightmarketplace/constants"
	"tui/flightmarketplace/errors"

	"github.com/gorilla/handlers"
)

// log into the terminal all the informations about a call to an api
func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
func recoveryPanicMdlw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a  function that will wait for a panic to arise
		defer func() {

			// recover from the panic
			rec := recover()
			if rec != nil {

				// try to cast the error to the custom roomenta error
				err, ok := rec.(errors.MarktErr)
				if ok {
					// if successful, send the error parameters as a response
					log.Println("PANIC: ", err)
					sendCustomError(w, err.StatusCode(), err.Params())
				} else {
					// if not succesful, send a default error
					log.Println("PANIC: panic was not a custom roomenta error.")
					err := errors.ErrServerError("panic was not a custom roomenta error.")
					sendCustomError(w, err.StatusCode(), err.Params())
				}
			}
		}()

		h.ServeHTTP(w, r)
	})
}

// sendCustomError encodes a roomerror's parameters and sends it as a response
func sendCustomError(w http.ResponseWriter, code int, param map[string]interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(param)
	if err != nil {
		log.Println("ERROR: error when encoding JSON panic response.")
	}
}

// adminAuthMiddleware mdlw for admin group authorization.
func adminAuthMiddleware(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return authorizeIfUserInGroup([]string{constants.AdminGroupManage}, handle)
}

// masterAuthMiddleware mdlw for master & admin groups authorization.
func marketerAuthMiddleware(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return authorizeIfUserInGroup([]string{constants.MarketerGroupManage}, handle)
}

// authorizeIfUserInGroup authorize user if he belongs to ALL authorized groups.
func authorizeIfUserInGroup(authorizedGroups []string, handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	// if authentication is disabled
	if config.Config.DisableKeycloakAuthentication {
		return func(w http.ResponseWriter, r *http.Request) { handle(w, r) }
	}

	return func(w http.ResponseWriter, r *http.Request) {

		/* 		claims := utility.GetUserFromContext(r)
		   		authorizedGroupCount := 0

		   		for _, authorizedGroup := range authorizedGroups {
		   			for _, group := range claims.Groups {
		   				if group == authorizedGroup {
		   					authorizedGroupCount += 1
		   					break
		   				}
		   			}
		   		}

		   		if authorizedGroupCount > 0 {
		   			handle(w, r)
		   			return
		   		} */

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("The user is not in the requested group"))
	}
}
