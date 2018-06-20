package protocol_meter

import ()

type ToDTUReadFramePkg struct {
	Addr     string
	CtrlCode uint8
	Data     []byte
}

func (p *ToDTUReadFramePkg) Parse(b []byte) {
	addr, ctrl_code, data := parse(b)

	for i, _ := range data {
		data[i] -= PROTOCOL_METER_DATA_SALT_BYTE
	}

	p.Addr = addr
	p.CtrlCode = ctrl_code
	p.Data = data
}
