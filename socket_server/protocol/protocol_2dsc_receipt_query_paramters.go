package protocol

import ()

type ToDSCReceiptQueryParamtersPkg struct {
	ID  [11]byte
	Tlv *TLV
}

func Parse2DSCReceiptQueryParamters(p []byte) *ToDSCReceiptQueryParamtersPkg {
	reader, _, id := parse_header(p)
	tlv := read_tlv(reader)

	return &ToDSCReceiptQueryParamtersPkg{
		ID:  id,
		Tlv: tlv,
	}
}
