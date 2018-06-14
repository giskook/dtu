package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/conf"
	"log"
	"net"
	"sync"
	"time"
)

type SocketServer struct {
	conf      *conf.Conf
	srv       *gotcp.Server
	cm        *ConnMgr
	SocketIn  chan base.Proto
	SocketOut chan base.Proto
	exit      chan struct{}
	wait_exit *sync.WaitGroup
	conn_uuid uint32
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	return &SocketServer{
		conf:      conf,
		cm:        NewConnMgr(),
		SocketIn:  make(chan base.Proto),
		SocketOut: make(chan base.Proto),
		exit:      make(chan struct{}),
		wait_exit: new(sync.WaitGroup),
	}
}

func (ss *SocketServer) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+ss.conf.TcpServer.BindPort)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}

	ss.srv = gotcp.NewServer(config, ss, ss)

	go ss.srv.Start(listener, time.Second)
	log.Println("<INFO> socket listening:", listener.Addr())

	for i := 0; i < ss.conf.TcpServer.WorkerNum; i++ {
		ss.consumer_worker()
	}

	return nil
}

func (ss *SocketServer) Send(id [11]byte, p gotcp.Packet) error {
	c := ss.cm.Get(id)
	if c != nil && c.status == USER_STATUS_NORMAL {
		c.SetWriteDeadline()
		return c.Send(p)
	}

	return base.ERROR_DTU_OFFLINE

}

func (ss *SocketServer) Stop() {
	close(ss.exit)
	ss.wait_exit.Wait()
	close(ss.SocketOut)
	close(ss.SocketIn)

	ss.srv.Stop()
}

func (ss *SocketServer) SetStatus(id [11]byte, status uint8) *Connection {
	c := ss.cm.Get(id)
	if c != nil {
		c.status = status
	}

	return c
}
