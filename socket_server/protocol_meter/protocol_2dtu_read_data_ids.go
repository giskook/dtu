package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToDTUReadDataElectricityPkg struct {
	Electricity string
}

func (p *ToDTUReadDataElectricityPkg) Parse(r *bytes.Reader) {
	interger := base.ReadBcdString(r, 3)
	decimal := base.ReadBcdString(r, 1)
	p.Electricity = interger + "." + decimal
}
