package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToMeterWriteAddrPkg struct {
	Addr uint64
}

func (p *ToMeterWriteAddrPkg) Serialize() []byte {
	var w bytes.Buffer
	write_header(&w, PROTOCOL_METER_ADDR_WILDCARD, PROTOCOL_METER_CTRL_CODE_2METER_WRITE_ADDR)
	base.WriteByte(&w, 0)
	base.WriteBytes(&w, base.GetBytePlus33LFix6(p.Addr))
	write_tail(&w)

	return w.Bytes()
}
