package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"user-event-store/management"
)

var Session *gocql.Session

func init() {
	var err error

	configuration := management.Configuration.Config("cassandra")
	host := configuration.String("host")
	keyspace := configuration.String("keyspace")

	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
