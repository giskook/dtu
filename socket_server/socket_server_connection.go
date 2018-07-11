package socket_server

import (
	"bytes"
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/socket_server/protocol_meter"
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

type meter struct {
	data      []byte
	timestamp int64
}

type Connection struct {
	MID uint64

	conf          *ConnConf
	c             *gotcp.Conn
	dtu_id        [11]byte
	RecvBuffer    *bytes.Buffer
	status        uint8
	exit          chan struct{}
	ticker        *time.Ticker
	err_recv_time int

	meter_addr     string
	meter_data     map[uint32]*meter
	ChanMDI        chan uint32
	wait           uint8
	wait_status    uint32
	meter_yymmddww string
}

func NewConnection(c *gotcp.Conn, conf *ConnConf) *Connection {
	tcp_c := c.GetRawConn()
	tcp_c.SetReadDeadline(time.Now().Add(time.Duration(conf.read_limit) * time.Second))
	tcp_c.SetWriteDeadline(time.Now().Add(time.Duration(conf.write_limit) * time.Second))
	return &Connection{
		conf:        conf,
		c:           c,
		RecvBuffer:  bytes.NewBuffer([]byte{}),
		meter_addr:  protocol_meter.PROTOCOL_METER_ADDR_WILDCARD,
		exit:        make(chan struct{}),
		ticker:      time.NewTicker(time.Duration(conf.interval) * time.Second),
		meter_data:  make(map[uint32]*meter),
		ChanMDI:     make(chan uint32),
		wait_status: protocol_meter.PROTOCOL_METER_DATA_ID_INVALID,
	}
}

func (c *Connection) req(data_id uint32) {
	c.wait = 1
	c.wait_status = data_id
	select {
	case <-c.ChanMDI:
		log.Println("<INF> recv end")
	case <-time.After(3 * time.Second):
		log.Println("<INF> recv timeout")
	}
	c.wait = 0
	c.wait_status = protocol_meter.PROTOCOL_METER_DATA_ID_INVALID
}

func (c *Connection) run(data_id uint32) {
	if data_id == c.wait_status &&
		c.wait == 1 {
		c.ChanMDI <- 0
		c.wait = 0
		c.wait_status = protocol_meter.PROTOCOL_METER_DATA_ID_INVALID
	}

}

func (c *Connection) do() {
	defer func() {
		c.c.Close()
	}()

	for {
		select {
		case <-c.exit:
			return
		case <-c.ticker.C:
			if c.status >= DTU_STATUS_METER_REG {
				c.meter_read_va()
				c.meter_read_a()
				c.meter_read_freeze_one()
				c.meter_read_freeze_two()
				c.meter_read_freeze_three()
				c.meter_read_freeze_four()
				c.meter_read_freeze_five()
				c.meter_read_freeze_six()
				c.meter_read_freeze_one_time()
				c.meter_read_freeze_two_time()
				c.meter_read_freeze_three_time()
				c.meter_read_freeze_four_time()
				c.meter_read_freeze_five_time()
				c.meter_read_freeze_six_time()
				c.meter_read_freeze_seven()
				c.meter_read_freeze_seven_time()
				c.meter_read_combine_elec_last_one()
				c.meter_read_combine_elec_last_two()
				c.meter_read_combine_elec_last_three()
				//				c.meter_read_positive_elec_last_one()
				//				c.meter_read_positive_elec_last_two()
				//				c.meter_read_positive_elec_last_three()
				c.meter_read_yymmddww()
				c.meter_read_hhmmss()
				c.meter_read_electricity()
			}
		}
	}
}

func (c *Connection) SetReadDeadline() {
	c.c.GetRawConn().SetReadDeadline(time.Now().Add(time.Duration(c.conf.read_limit) * time.Second))
}

func (c *Connection) SetWriteDeadline() {
	c.c.GetRawConn().SetWriteDeadline(time.Now().Add(time.Duration(c.conf.write_limit) * time.Second))
}

func (c *Connection) CloseSocket() {
	c.c.Close()
}

func (c *Connection) Close() {
	c.RecvBuffer.Reset()
	c.ticker.Stop()
	close(c.ChanMDI)
	close(c.exit)
}

func (c *Connection) Equal(cc *Connection) bool {
	return c.conf.uuid == cc.conf.uuid
}

func (c *Connection) Send(p gotcp.Packet) error {
	if c != nil && c.c != nil {
		log.Printf("<OUT> %x\n", p.Serialize())
		c.SetWriteDeadline()
		return c.c.AsyncWritePacket(p, 0)
	}

	return base.ERROR_DTU_OFFLINE
}
