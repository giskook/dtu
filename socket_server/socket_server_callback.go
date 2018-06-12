package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/socket_server/protocol"
	"log"
	//"runtime/debug"
	"sync/atomic"
)

func (ss *SocketServer) OnConnect(c *gotcp.Conn) bool {
	connection := NewConnection(c, &ConnConf{
		read_limit:  ss.conf.TcpServer.ReadLimit,
		write_limit: ss.conf.TcpServer.WriteLimit,
		uuid:        atomic.AddUint32(&ss.conn_uuid, 1),
	})

	c.PutExtraData(connection)
	//go connection.Check()
	log.Printf("<CNT> %x \n", c.GetRawConn())

	return true
}

func (ss *SocketServer) OnClose(c *gotcp.Conn) {
	connection := c.GetExtraData().(*Connection)
	ss.cm.Del(connection)
	connection.Close()
	log.Printf("<DIS> %x\n", c.GetRawConn())
	//debug.PrintStack()
}

func (ss *SocketServer) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	connection := c.GetExtraData().(*Connection)
	connection.SetReadDeadline()
	connection.RecvBuffer.Write(p.Serialize())
	for {
		protocol_id, protocol_length := protocol.CheckProtocol(connection.RecvBuffer)
		log.Printf("%d %d\n", protocol_id, protocol_length)
		buf := make([]byte, protocol_length)
		connection.RecvBuffer.Read(buf)
		if protocol_id != protocol.PROTOCOL_2DSC_REGISTER && connection != nil && connection.status != USER_STATUS_NORMAL {
			log.Printf("<SWALLOW> %x %x ", connection.ID, buf)
			return true
		}
		switch protocol_id {
		case protocol.PROTOCOL_HALF_PACK:
			return true
		case protocol.PROTOCOL_ILLEGAL:
			return true
		case protocol.PROTOCOL_2DSC_REGISTER:
			ss.eh_2dsc_register(buf, connection)
		case protocol.PROTOCOL_2DSC_DATA:
			ss.eh_2dsc_data(buf)
		}
	}
}
