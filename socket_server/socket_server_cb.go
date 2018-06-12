package socket_server

import (
	"github.com/giskook/dtu/socket_server/protocol"
)

func (ss *SocketServer) eh_2dsc_register(p []byte, c *Connection) {
	reg := protocol.Parse2DSCRegister(p)
	c.ID = reg.ID
	c.status = USER_STATUS_NORMAL
	ss.cm.Put(reg.ID, c)
	c.Send(reg)
	//ss.Send(reg.ID, reg)
}

func (ss *SocketServer) eh_2dsc_data(p []byte) {
	protocol.Parse2DSCData(p)
}
