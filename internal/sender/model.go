package sender

import (
	log "fmt"
	"net"
)

type ISender interface {
	ConnectToServer(string) *net.Conn
}

type Sender struct {
}

func (s *Sender) ConnectToServer(serverAdderss string) *net.Conn {

	connection, error := net.Dial("tcp4", serverAdderss)

	if error != nil {
		log.Println(error)
		return nil
	}

	return &connection

}
