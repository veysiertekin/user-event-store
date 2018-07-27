package telnet

import (
	"github.com/reiver/go-telnet"
	"user-event-store/app/utility/panic_helper"
	"log"
)

type Communicator struct {
	ServerAddress string
}

func (communicator Communicator) WriteString(message string) (err error) {
	return communicator.WriteByte([]byte(message))
}

func (communicator Communicator) WriteByte(message []byte) (err error) {
	defer panic_helper.RecoverOnError(&err)
	log.Printf("Message is currently being sent... <%s>", message)

	conn, err := telnet.DialTo(communicator.ServerAddress)
	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte("\n"))
	return err
}
