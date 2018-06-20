package socket_server

import (
	"bytes"
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/base"
	"log"
	"time"
)

const (
	DTU_STATUS_CNT       uint8 = 0
	DTU_STATUS_REG       uint8 = 1
	DTU_STATUS_METER_REG uint8 = 2
)

type ConnConf struct {
	read_limit  int
	write_limit int
	uuid        uint32
	interval    int
}

type Connection struct {
	conf       *ConnConf
	c          *gotcp.Conn
	ID         [11]byte
	RecvBuffer *bytes.Buffer
	status     uint8
	exit       chan struct{}
	ticker     *time.Ticker

	meter_addr string
}

func NewConnection(c *gotcp.Conn, conf *ConnConf) *Connection {
	tcp_c := c.GetRawConn()
	tcp_c.SetReadDeadline(time.Now().Add(time.Duration(conf.read_limit) * time.Second))
	tcp_c.SetWriteDeadline(time.Now().Add(time.Duration(conf.write_limit) * time.Second))
	return &Connection{
		conf:       conf,
		c:          c,
		RecvBuffer: bytes.NewBuffer([]byte{}),
		exit:       make(chan struct{}),
		ticker:     time.NewTicker(time.Duration(conf.interval) * time.Second),
	}
}

func (c *Connection) Do() {
	defer func() {
		c.c.Close()
	}()

	for {
		select {
		case <-c.exit:
			return
		case <-c.ticker.C:
		}
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
	c.ticker.Stop()
	close(c.exit)
}

func (c *Connection) Equal(cc *Connection) bool {
	return c.conf.uuid == cc.conf.uuid
}

func (c *Connection) Send(p gotcp.Packet) error {
	if c != nil && c.c != nil {
		log.Printf("<OUT> %x\n", p.Serialize())
		return c.c.AsyncWritePacket(p, 0)
	}

	return base.ERROR_DTU_OFFLINE
}
