package conf

import (
	"encoding/json"
	"os"
)

type TcpServer struct {
	BindPort          string
	ReadLimit         int
	WriteLimit        int
	ConnTimeout       int
	ConnCheckInterval int
	WorkerNum         int
	Interval          int
}

type HttpConf struct {
	Addr    string
	TimeOut int
}

type ZmqConf struct {
	TerminalManage2Dps string
	TerminalManage2Das string
	TerminalData2Dps   string
}

type Conf struct {
	TcpServer *TcpServer
	Http      *HttpConf
	Zmq       *ZmqConf
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
