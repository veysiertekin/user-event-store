package health_check

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/health").Subrouter()
	subRouter.Methods("GET").HandlerFunc(SendHeartbeat)
}
