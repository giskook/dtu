package reactor

import (
	"log"
)

func (rc *Reactor) shunt() {
	defer func() {
	}()

	for {
		select {
		case <-rc.exit:
			return
		case socket_in := <-rc.socket_server.SocketIn:
			log.Println(socket_in)
		}

	}
}
