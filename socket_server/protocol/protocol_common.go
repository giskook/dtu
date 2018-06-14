package protocol

import (
	"bytes"
	"encoding/binary"
	"github.com/giskook/dtu/base"
)

const (
	PROTOCOL_COMMON_LENGTH uint16 = 16
	PROTOCOL_MIN_LENGTH    uint16 = 16
	PROTOCOL_MAX_LENGTH    uint16 = 1046
	PROTOCOL_START_FLAG    uint8  = 0x7B
	PROTOCOL_END_FLAG      uint8  = 0x7B

	PROTOCOL_ILLEGAL   uint8 = 0xff
	PROTOCOL_HALF_PACK uint8 = 0xfe
	PROTOCOL_UNKNOWN   uint8 = 0

	PROTOCOL_2DSC_REGISTER                uint8 = 0x01
	PROTOCOL_2DSC_LOGOUT                  uint8 = 0x02
	PROTOCOL_2DSC_INVALID                 uint8 = 0x04
	PROTOCOL_2DSC_RECEIPT_DSC_DATA        uint8 = 0x05
	PROTOCOL_2DSC_DATA                    uint8 = 0x09
	PROTOCOL_2DSC_RECEIPT_QUERY_PARAMTERS uint8 = 0x0B
	PROTOCOL_2DSC_RECEIPT_SET_PARAMTERS   uint8 = 0x0D
	PROTOCOL_2DSC_RECEIPT_QUERY_LOG       uint8 = 0x0E
	PROTOCOL_2DSC_RECEIPT_UPGRADE         uint8 = 0x0F

	PROTOCOL_2DTU_RECEIPT_REGISTER uint8 = 0x81
	PROTOCOL_2DTU_RECEIPT_LOGOUT   uint8 = 0x82
	PROTOCOL_2DTU_REQ_REGISTER     uint8 = 0x83
	PROTOCOL_2DTU_INVALID          uint8 = 0x84
	PROTOCOL_2DTU_RECEIPT_DTU_DATA uint8 = 0x85
	PROTOCOL_2DTU_DATA             uint8 = 0x89
	PROTOCOL_2DTU_QUERY_PARAMTERS  uint8 = 0x8B
	PROTOCOL_2DTU_SET_PARAMTERS    uint8 = 0x8D
	PROTOCOL_2DTU_QUERY_LOG        uint8 = 0x8E
	PROTOCOL_2DTU_UPGRADE          uint8 = 0x8F

	PROTOCOL_2DTU_QUERY_PARAMTERS_ALL     uint8 = 0x00
	PROTOCOL_2DTU_QUERY_PARAMTERS_CMNET   uint8 = 0x01
	PROTOCOL_2DTU_QUERY_PARAMTERS_RTU     uint8 = 0x02
	PROTOCOL_2DTU_QUERY_PARAMTERS_SMS     uint8 = 0x03
	PROTOCOL_2DTU_QUERY_PARAMTERS_RUNTIME uint8 = 0x04
	PROTOCOL_2DTU_QUERY_PARAMTERS_SYS     uint8 = 0x05
	PROTOCOL_2DTU_QUERY_PARAMTERS_IP      uint8 = 0x06
)

type TLV struct {
	Type   uint8
	Flag   uint8
	Length uint16
	Value  []byte
}

func read_tlv(r *bytes.Reader) *TLV {
	_type := base.ReadByte(r)
	_flag := base.ReadByte(r)
	_length := base.ReadWord(r)
	_value := base.ReadBytes(r, int(_length))

	return &TLV{
		Type:   _type,
		Flag:   _flag,
		Length: _length,
		Value:  _value,
	}
}

func write_header(writer *bytes.Buffer, cmdid uint8, id []byte) {
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteByte(writer, cmdid)
	base.WriteWord(writer, 0)
	base.WriteBytesLength(writer, id, 11)
}

func write_tail(writer *bytes.Buffer) {
	write_length(writer)
	base.WriteByte(writer, PROTOCOL_END_FLAG)
}

func write_length(writer *bytes.Buffer) {
	length := writer.Len()
	length += 1
	length_byte := writer.Bytes()[2:4]
	binary.BigEndian.PutUint16(length_byte, uint16(length))
}

func parse_header(buffer []byte) (*bytes.Reader, uint16, [11]byte) {
	reader := bytes.NewReader(buffer)
	reader.Seek(2, 0)
	length := base.ReadWord(reader)
	id := base.ReadBytes(reader, 11)

	var result [11]byte
	copy(result[:], id)

	return reader, length, result
}

func CheckProtocol(buffer *bytes.Buffer) (uint8, int) {
	cmd := PROTOCOL_ILLEGAL
	cmd_len := 0
	bufferlen := uint16(buffer.Len())
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, 0
	}
	if buffer.Bytes()[0] != PROTOCOL_START_FLAG {
		buffer.ReadByte()
		cmd, cmd_len = CheckProtocol(buffer)
	} else if bufferlen >= PROTOCOL_MIN_LENGTH {
		len_in_protocol := base.GetWord(buffer.Bytes()[2:4])
		if len_in_protocol < PROTOCOL_MIN_LENGTH || len_in_protocol > PROTOCOL_MAX_LENGTH {
			buffer.ReadByte()
			cmd, cmd_len = CheckProtocol(buffer)
		}
		if len_in_protocol > bufferlen {
			return PROTOCOL_HALF_PACK, 0
		} else {
			if buffer.Bytes()[len_in_protocol-1] == PROTOCOL_END_FLAG {

				cmd = buffer.Bytes()[1]
				cmd_len = int(len_in_protocol)
				return cmd, cmd_len
			} else {
				buffer.ReadByte()
				cmd, cmd_len = CheckProtocol(buffer)
			}
		}
	} else {
		return PROTOCOL_HALF_PACK, 0
	}

	return cmd, cmd_len
}
