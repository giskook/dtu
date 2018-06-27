package socket_server

import (
	"github.com/giskook/dtu/socket_server/protocol"
	"github.com/giskook/dtu/socket_server/protocol_meter"
	"log"
)

func (c *Connection) meter_get_addr() {
	ppp := &protocol_meter.ToMeterReadAddrPkg{
		Addr: protocol_meter.PROTOCOL_METER_ADDR_WILDCARD,
	}
	log.Printf("<MOA> %x %x\n", c.c.GetRawConn(), ppp.Serialize())

	c.meter_send_cmd(ppp)
}

func (c *Connection) meter_write_addr(addr uint64) {
	p := &protocol_meter.ToMeterWriteAddrPkg{
		Addr: addr,
	}
	log.Printf("<MOA> %x %x\n", c.c.GetRawConn(), p.Serialize())

	c.meter_send_cmd(p)

}

func (c *Connection) meter_send_cmd(p protocol_meter.Packet) {
	c.Send(&protocol.ToDTUDataPkg{
		ID:  c.dtu_id,
		Pkg: p,
	})
}

func (c *Connection) meter_send_cmd_read_data(data_id uint32) {
	p := &protocol_meter.ToMeterReadDataPkg{
		Addr:   c.meter_addr,
		DataID: data_id,
	}
	log.Printf("<MOD> %x %x\n", c.c.GetRawConn(), p.Serialize())
	c.meter_send_cmd(p)
}

func (c *Connection) meter_read_electricity() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY)
}

func (c *Connection) meter_read_no() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_NO)
}

func (c *Connection) meter_read_va() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_VA)
}

func (c *Connection) meter_read_freeze_one() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE)
}

func (c *Connection) meter_read_freeze_two() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO)
}

func (c *Connection) meter_read_freeze_three() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE)
}

func (c *Connection) meter_read_freeze_four() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR)
}

func (c *Connection) meter_read_freeze_five() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE)
}

func (c *Connection) meter_read_freeze_six() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX)
}

func (c *Connection) meter_read_freeze_seven() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN)
}
