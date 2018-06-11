package socket_server

import (
	"github.com/giskook/dtu/socket_server/protocol"
)

func (ss *SocketServer) eh_2dsc_register(p []byte, c *Connection) {
	reg := protocol.Parse2DSCRegister(p)
	c.ID = reg.ID
	ss.cm.Put(reg.ID, c)
	ss.Send(reg.ID, reg)
}
