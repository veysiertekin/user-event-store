package health_check

import (
	"github.com/gorilla/mux"
	"user-event-store/app/endpoint/health_check/controller"
)

func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/health").Subrouter()
	subRouter.Methods("GET").HandlerFunc(controller.SendHeartbeat)
}
