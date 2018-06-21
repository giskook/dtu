package socket_server

import (
	"github.com/giskook/dtu/socket_server/protocol"
	"github.com/giskook/dtu/socket_server/protocol_meter"
	"log"
)

func (c *Connection) meter_send_cmd(p protocol_meter.Packet) {
	c.Send(&protocol.ToDTUDataPkg{
		ID:  c.ID,
		Pkg: p,
	})
}

func (c *Connection) meter_send_cmd_read_data(data_id uint32) {
	c.meter_send_cmd(&protocol_meter.ToMeterReadDataPkg{
		Addr:   c.meter_addr,
		DataID: data_id,
	})
}

func (c *Connection) meter_get_addr() {
	ppp := &protocol_meter.ToMeterReadAddrPkg{
		Addr: protocol_meter.PROTOCOL_METER_ADDR_WILDCARD,
	}
	log.Printf("-------%x\n", ppp.Serialize())

	c.meter_send_cmd(ppp)
}

func (c *Connection) meter_read_electricity() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY)
}

func (c *Connection) meter_read_no() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_NO)
}
