package protocol

import (
	"bytes"
)

type ToDTUReqRegisterPkg struct {
	ID []byte
}

func (p *ToDTUReqRegisterPkg) Serialize() []byte {
	var writer bytes.Buffer
	write_header(&writer, PROTOCOL_2DTU_REQ_REGISTER, p.ID)
	write_tail(&writer)

	return writer.Bytes()
}
