package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToDTUReadDataElectricityPkg struct {
	Electricity string
	ElecB       []byte
}

func (p *ToDTUReadDataElectricityPkg) Parse(r *bytes.Reader) {
	//	decimal := base.ReadBcdStringR(r, 1)
	//	interger := base.ReadBcdStringR(r, 3)
	//	p.Electricity = interger + "." + decimal
	p.ElecB = base.ReadBytes(r, 4)
}

type ToDTUReadDataNoPkg struct {
	No  string
	NoB []byte
}

func (p *ToDTUReadDataNoPkg) Parse(r *bytes.Reader) {
	p.NoB, p.No = base.ReadBcdStringRawR(r, 6)
}

type ToDTUReadDataVAPkg struct {
	VA []byte
}

func (p *ToDTUReadDataVAPkg) Parse(r *bytes.Reader) {
	p.VA, _ = base.ReadBcdStringRawR(r, 2)
}

type ToDTUReadDataFreezePkg struct {
	Elec []byte
}

func (p *ToDTUReadDataFreezePkg) Parse(r *bytes.Reader) {
	p.Elec, _ = base.ReadBcdStringRawR(r, 4)
}
