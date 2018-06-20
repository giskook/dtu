package protocol

import (
	"bytes"
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/socket_server/protocol_meter"
)

type ToDTUDataPkg struct {
	ID  [11]byte
	Pkg protocol_meter.Packet
}

func (p *ToDTUDataPkg) Serialize() []byte {
	var writer bytes.Buffer
	write_header(&writer, PROTOCOL_2DTU_DATA, p.ID[:])
	base.WriteBytes(&writer, p.Pkg.Serialize())
	write_tail(&writer)

	return writer.Bytes()
}
