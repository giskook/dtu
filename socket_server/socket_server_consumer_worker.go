package socket_server

import (
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/socket_server/protocol"
)

func (ss *SocketServer) consumer_worker() {
	ss.wait_exit.Add(1)
	go func() {
		for {
			select {
			case <-ss.exit:
				ss.wait_exit.Done()
				return
			case p := <-ss.SocketIn:
				http_type, id := p.Base()
				switch http_type {
				case base.PROTOCOL_2DTU_REQ_REGISTER:
					ss.Send(id, &protocol.ToDTUReqRegisterPkg{
						ID: id[:],
					})
				}
			}
		}
	}()
}
