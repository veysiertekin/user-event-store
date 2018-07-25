package telnet

import (
	"github.com/reiver/go-telnet"
	"user-event-store/utility/panic_helper"
	"log"
)

type Communicator struct {
	ServerAddress string // "host:port"
}

func (communicator Communicator) WriteMessage(message string) (err error) {
	defer panic_helper.RecoverOnError(&err)
	log.Printf("Message is currently being sent... <%s>", message)

	conn, err := telnet.DialTo(communicator.ServerAddress)
	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte("\n"))
	return err
}
