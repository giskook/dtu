package base

type Header struct {
	AppID string
	From  string
	To    string
}

type SocketData struct {
	Header *Header
	Data   []byte
}
