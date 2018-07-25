package user_event

import (
	"github.com/gorilla/mux"
	"user-event-store/app/endpoint/user_event/controller"
)

func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/user-event").Subrouter()
	subRouter.Methods("POST").HandlerFunc(controller.CreateUserEvent)
}
