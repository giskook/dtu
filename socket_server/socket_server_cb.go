package socket_server

import (
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/pb"
	"github.com/giskook/dtu/socket_server/protocol"
	"github.com/giskook/dtu/socket_server/protocol_meter"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
	"time"
)

func (ss *SocketServer) eh_2dsc_register(p []byte, c *Connection) {
	reg := protocol.Parse2DSCRegister(p)
	c.dtu_id = reg.ID

	if c.status < DTU_STATUS_REG {
		c.status = DTU_STATUS_REG
	}
	//ss.cm.Put(reg.ID, c)
	c.Send(reg)
	//c.meter_get_addr()
	if c.status == DTU_STATUS_REG {
		c.meter_read_no()
	}
}

func (ss *SocketServer) eh_2dsc_data(p []byte, c *Connection) {
	pp := protocol.Parse2DSCData(p)
	var frame protocol_meter.ToDTUReadFramePkg
	frame.Parse(pp.Data)
	log.Printf("<MI>  %x %x ctrl code %d\n", c.c.GetRawConn(), pp.Data, frame.CtrlCode)

	mid, _ := strconv.ParseUint(frame.Addr, 10, 64)
	c.MID = mid

	switch frame.CtrlCode {
	case protocol_meter.PROTOCOL_METER_CTRL_CODE_2DTU_READ_ADDR:
		var to_dtu_read_data_addr protocol_meter.ToDTUReadDataAddrPkg
		to_dtu_read_data_addr.Parse(frame.Data)
		c.meter_addr = to_dtu_read_data_addr.Addr
		c.err_recv_time = 0
		break
	case protocol_meter.PROTOCOL_METER_CTRL_CODE_2DTU_READ_DATA:
		ss.eh_2dsc_data_2dtu_read_data(frame.Data, c)
		c.err_recv_time = 0
		break
	case protocol_meter.PROTOCOL_METER_CTRL_CODE_2DTU_WRITE_ADDR:
		//	mid, _ := strconv.ParseUint(frame.Addr, 10, 64)
		//	c.MID = mid
		c.status = DTU_STATUS_METER_REG
		c.err_recv_time = 0

		break
	default:
		log.Printf("<INFO> ss eh_2dsc_data uncaught ctrl code %x\n", frame.CtrlCode)
		c.err_recv_time++
		if c.err_recv_time >= 4 {
			c.CloseSocket()
		}

	}

}

func (ss *SocketServer) eh_2dsc_data_2dtu_read_data_2dps_register(addr []byte, c *Connection) []byte {
	r := &Report.ManageCommand{
		Cpuid: addr,
		Type:  Report.ManageCommand_CMT_REQ_REGISTER,
		Uuid:  ss.conf.UUID,
		Tid:   uint64(c.conf.uuid),
		Paras: []*Report.Param{
			&Report.Param{},
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(c.conf.uuid),
			},
		},
	}
	data, _ := proto.Marshal(r)

	return data

}

func (ss *SocketServer) eh_2dsc_data_2dtu_read_data_2dps(d []byte, modbus_addr uint32, c *Connection) []byte {
	r := &Report.DataCommand{
		Uuid: ss.conf.UUID,
		//	Tid:  c.MID,
		Tid:  4,
		Type: Report.DataCommand_CMT_REP_DATA_UPLOAD_MONITORS,
		Monitors: []*Report.Monitor{
			&Report.Monitor{
				ModusAddr: modbus_addr,
				DataType:  1,
				DataLen:   uint32(len(d)) / 2,
				Data:      d,
			},
		},
	}
	data, _ := proto.Marshal(r)

	return data

}

func (ss *SocketServer) eh_2dsc_data_2dtu_read_data_all(c *Connection) []byte {
	ms := make([]*Report.Monitor, 0)
	for i, v := range c.meter_data {
		if len(v.data) == 4 {
			ms = append(ms, &Report.Monitor{
				ModusAddr: i,
				DataType:  1,
				DataLen:   2,
				Data:      v.data,
			})
		} else {
			ms = append(ms, &Report.Monitor{
				ModusAddr: i,
				DataType:  1,
				DataLen:   2,
				Data:      v.data,
			})
		}
	}
	r := &Report.DataCommand{
		Uuid: ss.conf.UUID,
		//	Tid:  c.MID,
		Tid:      1010001,
		Type:     Report.DataCommand_CMT_REP_DATA_UPLOAD_MONITORS,
		Monitors: ms,
	}
	data, _ := proto.Marshal(r)

	return data

}

var data_id_2register_id map[uint32]uint32 = map[uint32]uint32{protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_ONE: 19000, protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_TWO: 20000, protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_THREE: 21000}

func (ss *SocketServer) eh_2dsc_data_2dtu_read_data(b []byte, c *Connection) {
	r, data_id := protocol_meter.ParseDataID(b)
	log.Printf("<INF> recv data_id %x\n", data_id)
	switch data_id {
	case protocol_meter.PROTOCOL_METER_DATA_ID_NO:
		var e protocol_meter.ToDTUReadDataNoPkg
		e.Parse(r)
		log.Println("<INF> register to dps")
		ss.Socket2dps <- ss.eh_2dsc_data_2dtu_read_data_2dps_register(e.NoB, c)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_READ_ELECTRICITY:
		var elec protocol_meter.ToDTUReadDataElectricityPkg
		elec.Parse(r)
		c.meter_data[0x000003e8] = &meter{
			data:      base.C2B4(elec.Electricity, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		ss.Socket2dpsD <- ss.eh_2dsc_data_2dtu_read_data_all(c)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_V:
		var v protocol_meter.ToDTUReadDataVAPkg
		v.Parse(r)
		c.meter_data[0x000007d0] = &meter{
			data:      base.C2B4(v.VA, 10),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_A:
		var v protocol_meter.ToDTUReadDataAPkg
		v.Parse(r)
		c.meter_data[0x00000bb8] = &meter{
			data:      base.C2B4(v.A, 10),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00001770] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00001f40] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00002710] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00002ee0] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x000036b0] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00003e80] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN:
		var ef protocol_meter.ToDTUReadDataFreezePkg
		ef.Parse(r)
		c.meter_data[0x00004650] = &meter{
			data:      base.C2B4(ef.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_ONE_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00001388] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_TWO_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00001b58] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_THREE_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00002328] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FOUR_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00002af8] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_FIVE_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x000032c8] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SIX_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00003a98] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_FREEZE_SEVEN_TIME:
		var ef protocol_meter.ToDTUReadDataFreezeTimePkg
		ef.Parse(r)
		c.meter_data[0x00004268] = &meter{
			data:      base.C2B4(ef.TimeStamp, 1),
			timestamp: time.Now().Unix(),
		}
		log.Println(ef.TimeStamp)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_ONE, protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_TWO, protocol_meter.PROTOCOL_METER_DATA_ID_COMBINE_ELEC_LAST_THREE, protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_ONE, protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_TWO, protocol_meter.PROTOCOL_METER_DATA_ID_POSITIVE_ELEC_LAST_THREE:
		var ts protocol_meter.ToDTUReadDataSettlementPkg
		ts.Parse(r)
		c.meter_data[data_id_2register_id[data_id]] = &meter{
			data:      base.C2B4(ts.Elec, 100),
			timestamp: time.Now().Unix(),
		}
		log.Println(ts.Elec)
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_YYMMDDWW:
		var packet protocol_meter.ToDTUReadDataYYMMDDWWPkg
		packet.Parse(r)
		log.Println(packet.YYMMDDWW)
		c.meter_yymmddww = packet.YYMMDDWW
		c.run(data_id)
		break
	case protocol_meter.PROTOCOL_METER_DATA_ID_HHMMSS:
		var packet protocol_meter.ToDTUReadDataHHMMSSPkg
		packet.Parse(r)
		log.Println(packet.HHMMSS)
		c.meter_data[0x00000fa0] = &meter{
			data:      base.C2B4(string(c.meter_yymmddww[0:4])+packet.HHMMSS, 1),
			timestamp: time.Now().Unix(),
		}
		c.run(data_id)
		break

	default:
		log.Printf("ss eh_2dsc_data_2dtu_read_data uncaught data id %x\n", data_id)
	}
}
