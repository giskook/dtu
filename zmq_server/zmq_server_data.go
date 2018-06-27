package zmq_server

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func (z *ZmqServer) do_data_2dps() {
	for {
		select {
		case <-z.exit:
			return
		case p := <-z.ChanSTD2DPS:
			z.socket_terminal_data_2dps.SendBytes(p, zmq.DONTWAIT)
			log.Println("<INF> send data 2dps")
		}
	}
}
