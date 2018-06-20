package protocol_meter

import (
	"github.com/giskook/dtu/base"
)

type ToDTUReadDataAddrPkg struct {
	Addr string
}

func (p *ToDTUReadDataAddrPkg) Parse(b []byte) {
	p.Addr = base.GetBcdString(b)
}
