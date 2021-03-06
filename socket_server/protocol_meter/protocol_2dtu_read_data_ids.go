package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
	"log"
)

type ToDTUReadDataElectricityPkg struct {
	Electricity string
}

func (p *ToDTUReadDataElectricityPkg) Parse(r *bytes.Reader) {
	decimal := base.ReadBcdStringR(r, 1)
	interger := base.ReadBcdStringR(r, 3)
	p.Electricity = interger + "." + decimal
	log.Printf("<INF> Electricity %s\n", p.Electricity)
}

type ToDTUReadDataNoPkg struct {
	No  string
	NoB []byte
}

func (p *ToDTUReadDataNoPkg) Parse(r *bytes.Reader) {
	p.NoB, p.No = base.ReadBcdStringRawR(r, 6)
}

type ToDTUReadDataVAPkg struct {
	VA string
}

func (p *ToDTUReadDataVAPkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 2) // 3412
	p.VA = string(tmp[2]) + string(tmp[3]) + string(tmp[0]) + "." + string(tmp[1])
	log.Printf("<INF> VA %s\n", p.VA)
}

type ToDTUReadDataAPkg struct {
	A string
}

func (p *ToDTUReadDataAPkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 3) // 563412
	p.A = string(tmp[4]) + string(tmp[5]) + string(tmp[2]) + "." + string(tmp[3]) + string(tmp[0]) + string(tmp[1])
	log.Printf("<INF> A %s\n", p.A)
}

type ToDTUReadDataFreezePkg struct {
	Elec string
}

func (p *ToDTUReadDataFreezePkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 4) // 78563412
	p.Elec = string(tmp[6]) + string(tmp[7]) + string(tmp[4]) + string(tmp[5]) + string(tmp[2]) + string(tmp[3]) + "." + string(tmp[0]) + string(tmp[1])
	log.Printf("<INF> FE %s\n", p.Elec)
}

type ToDTUReadDataFreezeTimePkg struct {
	TimeStamp string
}

func (p *ToDTUReadDataFreezeTimePkg) Parse(r *bytes.Reader) {
	p.TimeStamp = base.ReadBcdString(r, 5)
	log.Printf("<INF> FT %s\n", p.TimeStamp)
}

type ToDTUReadDataSettlementPkg struct {
	Elec string
}

func (p *ToDTUReadDataSettlementPkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 4)
	p.Elec = string(tmp[6]) + string(tmp[7]) + string(tmp[4]) + string(tmp[5]) + string(tmp[2]) + string(tmp[3]) + "." + string(tmp[0]) + string(tmp[1])
	log.Printf("<INF> MF %s\n", p.Elec)
}

type ToDTUReadDataYYMMDDWWPkg struct {
	YYMMDDWW string
}

func (p *ToDTUReadDataYYMMDDWWPkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 4) //00080718

	p.YYMMDDWW = string(tmp[6]) + string(tmp[7]) + string(tmp[4]) + string(tmp[5]) + string(tmp[2]) + string(tmp[3])
	log.Printf("<INF> CTY %s\n", p.YYMMDDWW)
}

type ToDTUReadDataHHMMSSPkg struct {
	HHMMSS string
}

func (p *ToDTUReadDataHHMMSSPkg) Parse(r *bytes.Reader) {
	tmp := base.ReadBcdString(r, 3) // 344322

	p.HHMMSS = string(tmp[4]) + string(tmp[5]) + string(tmp[2]) + string(tmp[3]) + string(tmp[1]) + string(tmp[0])
	log.Printf("<INF> CTH %s\n", p.HHMMSS)
}
