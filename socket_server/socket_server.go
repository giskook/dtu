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
	conf        *conf.Conf
	srv         *gotcp.Server
	cm          *ConnMgr
	mcm         *ConnMgr
	Socket2das  chan []byte
	Socket2dps  chan []byte
	Socket2dpsD chan []byte
	exit        chan struct{}
	wait_exit   *sync.WaitGroup
	conn_uuid   uint32
}

func NewSocketServer(conf *conf.Conf) *SocketServer {
	return &SocketServer{
		conf:        conf,
		cm:          NewConnMgr(),
		mcm:         NewConnMgr(),
		Socket2das:  make(chan []byte),
		Socket2dps:  make(chan []byte),
		Socket2dpsD: make(chan []byte),
		exit:        make(chan struct{}),
		wait_exit:   new(sync.WaitGroup),
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

func (ss *SocketServer) Send(id uint64, p gotcp.Packet) error {
	c := ss.mcm.Get(id)
	if c != nil && c.status >= DTU_STATUS_CNT {
		c.SetWriteDeadline()
		return c.Send(p)
	}

	return base.ERROR_DTU_OFFLINE

}

func (ss *SocketServer) Stop() {
	close(ss.exit)
	ss.wait_exit.Wait()
	close(ss.Socket2dps)
	close(ss.Socket2das)
	close(ss.Socket2dpsD)

	ss.srv.Stop()
}

func (ss *SocketServer) SetStatus(id uint64, status uint8) *Connection {
	c := ss.mcm.Get(id)
	if c != nil {
		c.status = status
	}

	return c
}
