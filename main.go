package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"user-event-store/user_event"
	"user-event-store/health_check"
)

func main() {
	router := mux.NewRouter()
	user_event.RegisterRoutes(router)
	health_check.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
