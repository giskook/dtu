package protocol

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

type ToDSCDataPkg struct {
	ID   [11]byte
	Data []byte
}

func (p *ToDSCDataPkg) Serialize() []byte {
	var writer bytes.Buffer
	write_header(&writer, PROTOCOL_2DTU_RECEIPT_DTU_DATA, p.ID[:])
	write_tail(&writer)

	return writer.Bytes()
}

func Parse2DSCData(p []byte) *ToDSCDataPkg {
	reader, length, id := parse_header(p)
	data := base.ReadBytes(reader, int(length-PROTOCOL_COMMON_LENGTH))

	return &ToDSCDataPkg{
		ID:   id,
		Data: data,
	}
}
