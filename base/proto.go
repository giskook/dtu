package base

const (
	PROTOCOL_2DTU_REQ_REGISTER uint8 = 0x83
	PROTOCOL_2DSC_REGISTER           = 0x01
)

type InnerCmdDel struct {
	Type uint8
	ID   [11]byte
}

type HttpInOut struct {
	Req  Proto
	Resp chan Proto
}

type Proto interface {
	Base() (uint8, [11]byte)
}

type Restart struct {
	Type   uint8
	ID     [11]byte
	Result int
}

func (c *Restart) Base() (uint8, [11]byte) {
	return c.Type, c.ID
}

//////////////////////////////////////////////////////////////////

const (
	PROTOCOL_METER_2DSC_ELECTRICITY uint32 = 0x00000000
)

type ToDSCData interface {
	Base() (uint32, uint64)
}

type ToDSCDataElectricity struct {
	Mid         uint64
	Electricity string
}

func (e *ToDSCDataElectricity) Base() (uint32, uint64) {
	return PROTOCOL_METER_2DSC_ELECTRICITY, e.Mid
}
