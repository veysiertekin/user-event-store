package user_event

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/user-event").Subrouter()
	subRouter.Methods("POST").HandlerFunc(CreateUserEvent)
}
