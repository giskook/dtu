package http_srv

import (
	"github.com/giskook/dtu/base"
	"github.com/giskook/dtu/conf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf       *conf.Conf
	router     *mux.Router
	HttpInOut  chan *base.HttpInOut
	HttpCmdDel chan *base.InnerCmdDel
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	return &HttpSrv{
		conf:       conf,
		router:     mux.NewRouter(),
		HttpInOut:  make(chan *base.HttpInOut),
		HttpCmdDel: make(chan *base.InnerCmdDel),
	}
}

func (h *HttpSrv) Start() {
	s := h.router.PathPrefix("/plc").Subrouter()
	h.init_api_plc(s)

	if err := http.ListenAndServe(h.conf.Http.Addr, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("<INFO> http listening : %s\n", h.conf.Http.Addr)
}

func (h *HttpSrv) init_api_plc(r *mux.Router) {
	s := r.PathPrefix("/").Subrouter()
	s.HandleFunc("/restart", h.handler_dtu_restart)
}
