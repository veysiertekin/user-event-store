package main

import (
	"log"
	"fmt"
	"expvar"
	"net/http"
	"github.com/gorilla/mux"
	"user-event-store/app/client/cassandra"
	"user-event-store/app/endpoint/user_event"
	"user-event-store/app/endpoint/health_check"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(w, "%q: %q", key, str)
		} else {
			fmt.Fprintf(w, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}

func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter()
	router.HandleFunc("/debug/vars", metricsHandler)

	user_event.RegisterRoutes(router)
	health_check.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
