package zmq_server

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func (z *ZmqServer) do_manage_2dps() {
	for {
		select {
		case <-z.exit:
			return
		case p := <-z.ChanSTM2DPS:
			z.socket_terminal_manage_2dps.SendBytes(p, zmq.DONTWAIT)
		}
	}
}

func (z *ZmqServer) do_manage_2das() {
	for {
		select {
		case <-z.exit:
			return
		default:
			_, err := z.socket_terminal_manage_2das.Recv(0)
			if err != nil {
				log.Printf("<ERR> manage_2das %s\n", err.Error())
			}
			z.socket_terminal_manage_2das.Recv(0)
			z.socket_terminal_manage_2das.Recv(0)
			p, _ := z.socket_terminal_manage_2das.RecvBytes(0)
			z.ChanSTM2DAS <- p
		}
	}
}
