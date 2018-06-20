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
	log.Printf("-------%x\n", ppp.Serialize())
	c.Send(&protocol.ToDTUDataPkg{
		ID: c.ID,
		Pkg: &protocol_meter.ToMeterReadAddrPkg{
			Addr: protocol_meter.PROTOCOL_METER_ADDR_WILDCARD,
		},
	})
}

func (c *Connection) meter_read_electricity() {
	c.Send(&protocol.ToDTUDataPkg{
		ID: c.ID,
		Pkg: &protocol_meter.ToMeterReadDataPkg{
			Addr:   c.meter_addr,
			DataID: protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY,
		},
	})
}
