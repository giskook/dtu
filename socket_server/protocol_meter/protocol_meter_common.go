package protocol_meter

import (
	"bytes"
	"github.com/giskook/dtu/base"
)

const (
	PROTOCOL_METER_MIN_LENGTH           uint8 = 12
	PROTOCOL_METER_COMMON_HEADER_LENGTH uint8 = 10
	PROTOCOL_METER_START_FLAG           uint8 = 0x68
	PROTOCOL_METER_END_FLAG             uint8 = 0x16
	PROTOCOL_METER_ILLEGAL              uint8 = 0xff
	PROTOCOL_METER_HALF_PACK            uint8 = 0xfe
	PROTOCOL_METER_UNKNOW               uint8 = 0

	PROTOCOL_METER_CTRL_CODE_2METER_READ_DATA         uint8 = 0x11
	PROTOCOL_METER_CTRL_CODE_2DTU_READ_DATA           uint8 = 0x91
	PROTOCOL_METER_CTRL_CODE_2DTU_READ_DATA_MORE      uint8 = 0xB1
	PROTOCOL_METER_CTRL_CODE_2DTU_READ_DATA_EXCEPTION uint8 = 0xD1

	PROTOCOL_METER_CTRL_CODE_2METER_READ_ADDR uint8 = 0x13
	PROTOCOL_METER_CTRL_CODE_2DTU_READ_ADDR   uint8 = 0x93

	PROTOCOL_METER_CTRL_CODE_2METER_WRITE_ADDR uint8 = 0x15
	PROTOCOL_METER_CTRL_CODE_2DTU_WRITE_ADDR   uint8 = 0x95

	PROTOCOL_METER_DATA_ID_READ_ELECTRICITY uint32 = 0x00000000
	PROTOCOL_METER_DATA_ID_ADDR             uint32 = 0x04000401
	PROTOCOL_METER_DATA_ID_NO               uint32 = 0x04000402
	PROTOCOL_METER_DATA_ID_VA               uint32 = 0x02010100
	PROTOCOL_METER_DATA_ID_FREEZE_ONE       uint32 = 0x05000101
	PROTOCOL_METER_DATA_ID_FREEZE_TWO       uint32 = 0x05000102
	PROTOCOL_METER_DATA_ID_FREEZE_THREE     uint32 = 0x05000103
	PROTOCOL_METER_DATA_ID_FREEZE_FOUR      uint32 = 0x05000104
	PROTOCOL_METER_DATA_ID_FREEZE_FIVE      uint32 = 0x05000105
	PROTOCOL_METER_DATA_ID_FREEZE_SIX       uint32 = 0x05000106
	PROTOCOL_METER_DATA_ID_FREEZE_SEVEN     uint32 = 0x05000107

	PROTOCOL_METER_ADDR_WILDCARD   string = "AAAAAAAAAAAA"
	PROTOCOL_METER_DATA_SALT_BYTE  uint8  = 0x33
	PROTOCOL_METER_DATA_SALT_WORD  uint16 = 0x3333
	PROTOCOL_METER_DATA_SALT_DWORD uint32 = 0x33333333
)

func write_header(w *bytes.Buffer, addr string, ctrl_code uint8) {
	w.WriteByte(PROTOCOL_METER_START_FLAG)
	base.WriteBcdString(w, addr)
	base.WriteByte(w, PROTOCOL_METER_START_FLAG)
	base.WriteByte(w, ctrl_code)
}

func write_tail(w *bytes.Buffer) {
	w.Bytes()[9] = uint8(w.Len()) - PROTOCOL_METER_COMMON_HEADER_LENGTH
	base.WriteByte(w, sum(w.Bytes()))
	base.WriteByte(w, PROTOCOL_METER_END_FLAG)
}

func sum(d []byte) uint8 {
	var result uint8 = 0
	for _, v := range d {
		result += v
	}
	return result

}

func CheckProtocol(b []byte) (uint8, uint8) {
	cmd := PROTOCOL_METER_ILLEGAL
	cmd_len := uint8(0)
	bufferlen := uint8(len(b))
	if bufferlen == 0 {
		return PROTOCOL_METER_ILLEGAL, 0
	}
	if b[0] != PROTOCOL_METER_START_FLAG {
		cmd, cmd_len = CheckProtocol(b[1:])
	} else if bufferlen >= PROTOCOL_METER_MIN_LENGTH {
		len_in_protocol := b[9] + PROTOCOL_METER_MIN_LENGTH
		if len_in_protocol > bufferlen {
			return PROTOCOL_METER_HALF_PACK, 0
		} else {
			if b[len_in_protocol-1] == PROTOCOL_METER_END_FLAG &&
				b[len_in_protocol-2] == sum(b[0:len_in_protocol-2]) {

				cmd = b[8]
				cmd_len = len_in_protocol
				return cmd, cmd_len
			} else {
				cmd, cmd_len = CheckProtocol(b[1:])
			}
		}
	} else {
		return PROTOCOL_METER_HALF_PACK, 0
	}

	return cmd, cmd_len
}

type Packet interface {
	Serialize() []byte
}

func parse(buffer []byte) (string, uint8, []byte) {
	reader := bytes.NewReader(buffer)
	reader.Seek(1, 0)
	addr := base.ReadBcdStringR(reader, 6)
	base.ReadByte(reader)
	ctrl_code := base.ReadByte(reader)
	length := base.ReadByte(reader)
	data := base.ReadBytes(reader, int(length))

	return addr, ctrl_code, data
}
func ParseDataID(buffer []byte) (*bytes.Reader, uint32) {
	reader := bytes.NewReader(buffer)
	data_id := base.ReadDWordL(reader)

	return reader, data_id
}
