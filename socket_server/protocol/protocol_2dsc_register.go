package protocol

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToDSCRegisterPkg struct {
	ID   [11]byte
	IP   []byte
	Port uint16
}

func (p *ToDSCRegisterPkg) Serialize() []byte {
	var writer bytes.Buffer
	write_header(&writer, PROTOCOL_2DTU_RECEIPT_REGISTER)
	base.WriteWord(&writer, 0)
	base.WriteBytes(&writer, p.ID[:])
	write_tail(&writer)

	return writer.Bytes()
}

func Parse2DSCRegister(p []byte) *ToDSCRegisterPkg {
	reader, id := parse_header(p)
	ip := base.ReadBytes(reader, 4)
	port := base.ReadWord(reader)

	return &ToDSCRegisterPkg{
		ID:   id,
		IP:   ip,
		Port: port,
	}
}
