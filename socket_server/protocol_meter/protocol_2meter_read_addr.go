package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToMeterReadAddrPkg struct {
	Addr string
}

func (p *ToMeterReadAddrPkg) Serialize() []byte {
	var w bytes.Buffer
	write_header(&w, p.Addr, PROTOCOL_METER_CTRL_CODE_2METER_READ_ADDR)
	base.WriteByte(&w, 0)
	write_tail(&w)

	return w.Bytes()
}
