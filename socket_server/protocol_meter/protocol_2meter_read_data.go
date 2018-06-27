package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToMeterReadDataPkg struct {
	Addr   string
	DataID uint32
}

func (p *ToMeterReadDataPkg) Serialize() []byte {
	var w bytes.Buffer
	write_header(&w, p.Addr, PROTOCOL_METER_CTRL_CODE_2METER_READ_DATA)
	base.WriteByte(&w, 0)
	base.WriteDWordL(&w, p.DataID+PROTOCOL_METER_DATA_SALT_DWORD)
	write_tail(&w)

	return w.Bytes()
}
