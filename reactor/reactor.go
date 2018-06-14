package reactor

import (
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/conf"
	"github.com/giskook/dtu/http_srv"
	"github.com/giskook/dtu/socket_server"
	"log"
)

type Reactor struct {
	conf *conf.Conf
	exit chan struct{}

	socket_server *socket_server.SocketServer
	http_srv      *http_srv.HttpSrv
	hm            *HttpMgr
}

func NewReactor(conf *conf.Conf) *Reactor {
	return &Reactor{
		conf:          conf,
		exit:          make(chan struct{}),
		socket_server: socket_server.NewSocketServer(conf),
		http_srv:      http_srv.NewHttpSrv(conf),
		hm:            NewHttpMgr(),
	}
}

func (r *Reactor) Start() error {
	err := r.socket_server.Start()
	if err != nil {
		return err
	}

	go r.http_srv.Start()

	r.shunt()

	return nil
}

func (r *Reactor) Stop() {
	r.socket_server.Stop()
	log.Printf("<INFO> %s\n", base.SOCKET_SERVER_STOPPED)
	close(r.exit)
}
