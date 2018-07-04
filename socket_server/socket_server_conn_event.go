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

func (c *Connection) meter_read_no() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_NO)
}

func (c *Connection) meter_read_electricity() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY)
}

func (c *Connection) meter_read_va() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_V)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_V)

}

func (c *Connection) meter_read_a() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_A)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_A)
}

func (c *Connection) meter_read_freeze_one() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE)
}

func (c *Connection) meter_read_freeze_two() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO)
}

func (c *Connection) meter_read_freeze_three() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE)
}

func (c *Connection) meter_read_freeze_four() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR)
}

func (c *Connection) meter_read_freeze_five() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE)
}

func (c *Connection) meter_read_freeze_six() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX)
}

func (c *Connection) meter_read_freeze_seven() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN)
}

func (c *Connection) meter_read_freeze_one_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE_TIME)
}

func (c *Connection) meter_read_freeze_two_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO_TIME)
}

func (c *Connection) meter_read_freeze_three_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE_TIME)
}

func (c *Connection) meter_read_freeze_four_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR_TIME)
}

func (c *Connection) meter_read_freeze_five_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE_TIME)
}

func (c *Connection) meter_read_freeze_six_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX_TIME)
}

func (c *Connection) meter_read_freeze_seven_time() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN_TIME)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN_TIME)
}

func (c *Connection) meter_read_combine_elec_last_one() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_ONE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_ONE)
}

func (c *Connection) meter_read_combine_elec_last_two() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_TWO)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_TWO)
}

func (c *Connection) meter_read_combine_elec_last_three() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_THREE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_THREE)
}

func (c *Connection) meter_read_positive_elec_last_one() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_ONE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_ONE)
}

func (c *Connection) meter_read_positive_elec_last_two() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_TWO)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_TWO)
}

func (c *Connection) meter_read_positive_elec_last_three() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_THREE)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_THREE)
}

func (c *Connection) meter_read_yymmddww() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_YYMMDDWW)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_YYMMDDWW)
}

func (c *Connection) meter_read_hhmmss() {
	c.meter_send_cmd_read_data(protocol_meter.PROTOCOL_METER_DATA_ID_HHMMSS)
	c.req(protocol_meter.PROTOCOL_METER_DATA_ID_HHMMSS)
}
