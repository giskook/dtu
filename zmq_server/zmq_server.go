package zmq_server

import (
	"github.com/giskook/dtu/conf"
	//"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"log"
)

type ZmqServer struct {
	socket_terminal_manage_2dps *zmq.Socket
	socket_terminal_manage_2das *zmq.Socket
	ChanSTM2DPS                 chan string
	ChanSTM2DAS                 chan string

	socket_terminal_data_2dps *zmq.Socket
	ChanSTD2DPS               chan string

	exit chan struct{}
}

func NewZmqServer(cnf *conf.ZmqConf) *ZmqServer {
	stm2dps, _ := zmq.NewSocket(zmq.PUSH)
	err := stm2dps.Connect(cnf.TerminalManage2Dps)
	if err != nil {
		log.Printf("<ERR> %s\n", err.Error())
		return nil
	}

	stm2das, _ := zmq.NewSocket(zmq.SUB)
	err = stm2dps.Connect(cnf.TerminalManage2Das)
	if err != nil {
		log.Printf("<ERR> %s\n", err.Error())
		return nil
	}

	std2dps, _ := zmq.NewSocket(zmq.PUSH)
	err = std2dps.Connect(cnf.TerminalData2Dps)
	if err != nil {
		log.Printf("<ERR> %s\n", err.Error())
		return nil
	}

	return &ZmqServer{
		socket_terminal_manage_2dps: stm2dps,
		socket_terminal_manage_2das: stm2das,
		ChanSTM2DPS:                 make(chan string),
		ChanSTM2DAS:                 make(chan string),
		socket_terminal_data_2dps:   std2dps,
		ChanSTD2DPS:                 make(chan string),
		exit:                        make(chan struct{}),
	}
}
