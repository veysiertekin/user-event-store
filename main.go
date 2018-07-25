package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"user-event-store/app/client/cassandra"
	"user-event-store/app/endpoint/user_event"
	"user-event-store/app/endpoint/health_check"
)

func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter()
	user_event.RegisterRoutes(router)
	health_check.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
