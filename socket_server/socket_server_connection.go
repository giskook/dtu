package socket_server

import (
	"bytes"
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/base"
	"log"
	"time"
)

const (
	USER_STATUS_INIT    uint8 = 0
	USER_STATUS_NORMAL  uint8 = 1
	USER_STATUS_ILLEGAL uint8 = 2
)

type ConnConf struct {
	read_limit  int
	write_limit int
	uuid        uint32
}

type Connection struct {
	conf       *ConnConf
	c          *gotcp.Conn
	ID         [11]byte
	RecvBuffer *bytes.Buffer
	status     uint8
}

func NewConnection(c *gotcp.Conn, conf *ConnConf) *Connection {
	tcp_c := c.GetRawConn()
	tcp_c.SetReadDeadline(time.Now().Add(time.Duration(conf.read_limit) * time.Second))
	tcp_c.SetWriteDeadline(time.Now().Add(time.Duration(conf.write_limit) * time.Second))
	return &Connection{
		conf:       conf,
		c:          c,
		RecvBuffer: bytes.NewBuffer([]byte{}),
	}
}

func (c *Connection) SetReadDeadline() {
	c.c.GetRawConn().SetReadDeadline(time.Now().Add(time.Duration(c.conf.read_limit) * time.Second))
}

func (c *Connection) SetWriteDeadline() {
	c.c.GetRawConn().SetWriteDeadline(time.Now().Add(time.Duration(c.conf.write_limit) * time.Second))
}

func (c *Connection) Close() {
	c.RecvBuffer.Reset()
}

func (c *Connection) Equal(cc *Connection) bool {
	return c.conf.uuid == cc.conf.uuid
}

func (c *Connection) Send(p gotcp.Packet) error {
	if c != nil && c.c != nil {
		log.Printf("<OUT> %x\n", p.Serialize())
		return c.c.AsyncWritePacket(p, 0)
	}

	return base.ErrSocketAlreadyNotExist
}
