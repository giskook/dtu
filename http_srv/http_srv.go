package http_srv

import (
	"github.com/giskook/dtu/conf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf   *conf.Conf
	router *mux.Router
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	return &HttpSrv{
		conf:   conf,
		router: mux.NewRouter(),
	}
}

func (h *HttpSrv) Start() {
	s := h.router.PathPrefix("/plc").Subrouter()
	h.init_api_plc(s)

	if err := http.ListenAndServe(h.conf.Http.Addr, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h *HttpSrv) init_api_plc(r *mux.Router) {
	s := r.PathPrefix("/").Subrouter()
	s.HandleFunc("/restart", h.handler_dtu_restart)
}
