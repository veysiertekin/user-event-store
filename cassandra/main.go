package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"user-event-store/configuration"
)

var Session *gocql.Session

func init() {
	var err error

	config := configuration.Data.Config("cassandra")
	host := config.String("host")
	keyspace := config.String("keyspace")

	cluster := gocql.NewCluster(host)
	cluster.Keyspace = configuration.Data.String(keyspace)
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
