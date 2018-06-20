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

type Conf struct {
	TcpServer *TcpServer
	Http      *HttpConf
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
