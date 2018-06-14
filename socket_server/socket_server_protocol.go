package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/base"
	"log"
	"net"
)

var ()

type Raw struct {
	raw []byte
}

func (r *Raw) Serialize() []byte {
	return r.raw
}

func (ss *SocketServer) ReadPacket(conn *net.TCPConn) (gotcp.Packet, error) {
	data := make([]byte, 1024)
	length, err := conn.Read(data)
	if err != nil {
		log.Println("<ERR> %s\n", err.Error())
		return nil, err
	}

	if length == 0 {
		log.Println("<ERR> peer error\n")
		return nil, base.ERROR_DTU_CLOSE_CONN
	}
	log.Printf("<IN>  %x  %x\n", conn, data[0:length])

	return &Raw{
		raw: data[0:length],
	}, nil
}
