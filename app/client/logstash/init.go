package logstash

import (
	"fmt"
	"log"
	"user-event-store/app/management"
	"user-event-store/app/client/telnet"
	"encoding/json"
)

var Client *ClientWrapper

type ClientWrapper struct {
	Communicator telnet.Communicator
}

func (client ClientWrapper) SendJsonAsync(data interface{}) {
	message, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error while encoding:", err)
	}
	go client.Communicator.WriteByte(message)
}

func init() {
	configuration := management.Configuration.Config("logstash")

	host := configuration.String("Host")
	port := configuration.Int("Port")

	serverAddress := fmt.Sprintf("%s:%d", host, port)
	Client = &ClientWrapper{telnet.Communicator{ServerAddress: serverAddress}}

	log.Printf("Logstash client initialized, %v", Client)
}
