package socket_server

import (
	//"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/pb"
	//"github.com/giskook/dtu/socket_server/protocol"
	"github.com/golang/protobuf/proto"
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
			case p := <-ss.Socket2das:
				command := &Report.ManageCommand{}
				err := proto.Unmarshal([]byte(p), command)
				if err != nil {
					log.Println("<ERR> socket_server Socket2das unmarshal error")
				} else {
					switch command.Type {
					case Report.ManageCommand_CMT_REP_REGISTER:
						if c := ss.cm.Get(command.Paras[0].Npara); c != nil {
							if c.MID != command.Paras[1].Npara {
								c.meter_write_addr(command.Paras[1].Npara)
							} else {
								c.status = DTU_STATUS_METER_REG
							}
						}
					}
				}
			}
		}
	}()
}
