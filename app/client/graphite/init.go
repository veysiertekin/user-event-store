package graphite

import (
	"fmt"
	"log"
	"time"
	"user-event-store/app/management"
	"user-event-store/app/utility/telnet"
)

var Client *ClientWrapper

type ClientWrapper struct {
	Communicator telnet.Communicator
}

func (client ClientWrapper) SendMetricAsync(key string, value string) {
	message := fmt.Sprintf("%s %s %d", key, value, time.Now().Unix())
	go client.Communicator.WriteString(message)
}

func init() {
	configuration := management.Configuration.Config("graphite")

	host := configuration.String("Host")
	port := configuration.Int("Port")

	serverAddress := fmt.Sprintf("%s:%d", host, port)
	Client = &ClientWrapper{telnet.Communicator{ServerAddress: serverAddress}}

	log.Printf("Graphite client initialized, %v", Client)
}
