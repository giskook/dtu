package socket_server

import (
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/socket_server/protocol"
	"github.com/giskook/dtu/socket_server/protocol_meter"
	"log"
)

func (ss *SocketServer) eh_2dsc_register(p []byte, c *Connection) {
	reg := protocol.Parse2DSCRegister(p)
	c.ID = reg.ID
	c.status = DTU_STATUS_REG
	ss.cm.Put(reg.ID, c)
	c.Send(reg)
	//ss.Send(reg.ID, reg)
	ss.SocketOut <- &base.Restart{
		Type: base.PROTOCOL_2DSC_REGISTER,
		ID:   reg.ID,
	}
	c.meter_get_addr()
}

func (ss *SocketServer) eh_2dsc_data(p []byte, c *Connection) {
	pp := protocol.Parse2DSCData(p)
	var frame protocol_meter.ToDTUReadFramePkg
	frame.Parse(pp.Data)

	switch frame.CtrlCode {
	case protocol_meter.PROTOCOL_METER_CTRL_CODE_2DTU_READ_ADDR:
		var to_dtu_read_data_addr protocol_meter.ToDTUReadDataAddrPkg
		to_dtu_read_data_addr.Parse(frame.Data)
		c.meter_addr = to_dtu_read_data_addr.Addr
		c.status = DTU_STATUS_METER_REG
		break
	case protocol_meter.PROTOCOL_METER_CTRL_CODE_2DTU_READ_DATA:
		ss.eh_2dsc_data_2dtu_read_data(frame.Data)
		break
	default:
		log.Printf("ss eh_2dsc_data uncaught ctrl code %d\n", frame.CtrlCode)

	}

}

func (ss *SocketServer) eh_2dsc_data_2dtu_read_data(b []byte) {
	r, data_id := protocol_meter.ParseDataID(b)
	switch data_id {
	case protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY:
		var e protocol_meter.ToDTUReadDataElectricityPkg
		e.Parse(r)
		break
	default:
		log.Printf("ss eh_2dsc_data_2dtu_read_data uncaught data id %d\n", data_id)
	}
}
