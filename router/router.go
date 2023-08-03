package router

import (
	"log"
	"net/http"
	"tui/flightmarketplace/controller"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// SetupRouter setup the routing of roomenta
func SetupRouter(router *mux.Router, ctrl controller.ControllerInterface) {

	setupPingRouter(router)

	router.Use(loggingMiddleware, recoveryPanicMdlw)

	setupPrivateRouter(router.PathPrefix("").Subrouter(), ctrl)
	setupPublicRouter(router.PathPrefix("").Subrouter(), ctrl)
}

func setupPingRouter(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Pong"))
		if err != nil {
			log.Println("Server was unable to Pong in response to a Ping.")
		}
	})
}

// setup all the public routes, they don't need authentication
func setupPublicRouter(router *mux.Router, ctrl controller.ControllerInterface) {
	setupSwaggoDocumentation(router)
}

// setup all the private routes, they need authentication
func setupPrivateRouter(router *mux.Router, ctrl controller.ControllerInterface) {
	//router.Use(authenticationMiddleware)

	setupFlightRouter(router.PathPrefix("/flight").Subrouter(), ctrl)
	//setupOrderRouter(router.PathPrefix("/order").Subrouter(), ctrl)
	//setupUserRouter(router.PathPrefix("/user").Subrouter(), ctrl)
}
func setupSwaggoDocumentation(router *mux.Router) {
	// Necessary for swaggo (documentation)
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
}
