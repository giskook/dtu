package socket_server

import (
	"log"
)

func (ss *SocketServer) consumer_worker() {
	ss.wait_exit.Add(1)
	go func() {
		for {
			select {
			case <-ss.exit:
				ss.wait_exit.Done()
				return
			case p := <-ss.SocketOut:
				log.Println(p)
			}
		}
	}()
}
