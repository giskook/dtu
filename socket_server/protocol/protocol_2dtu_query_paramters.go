package protocol

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToDTUQueryParamtersPkg struct {
	ID     [11]byte
	Type   uint8
	Serial []byte
}

func (p *ToDTUQueryParamtersPkg) Serialize() []byte {
	var writer bytes.Buffer
	write_header(&writer, PROTOCOL_2DTU_QUERY_PARAMTERS, p.ID[:])
	base.WriteByte(&writer, p.Type)
	base.WriteBytes(&writer, p.Serial)
	write_tail(&writer)

	return writer.Bytes()
}
