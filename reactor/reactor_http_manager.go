package reactor

import (
	"github.com/giskook/dtu/base"
	"log"
	"sync"
)

type HttpMgr struct {
	requests map[[11]byte][]*base.HttpInOut
	mutex    *sync.RWMutex
}

func NewHttpMgr() *HttpMgr {
	return &HttpMgr{
		requests: make(map[[11]byte][]*base.HttpInOut),
		mutex:    new(sync.RWMutex),
	}
}

func (hm *HttpMgr) Put(id [11]byte, http_in_out *base.HttpInOut) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()

	if _, ok := hm.requests[id]; ok {
		hm.requests[id] = append(hm.requests[id], http_in_out)
		return
	}
	hm.requests[id] = []*base.HttpInOut{
		http_in_out,
	}
}

func (hm *HttpMgr) Get(id [11]byte, action uint8) []*base.HttpInOut {
	log.Println("---------------")
	log.Printf("dtu id %s  action %d \n", string(id[:]), action)
	log.Println(hm.requests)
	for _, kk := range hm.requests[id] {
		tt, idd := kk.Req.Base()
		log.Printf("hm  %s %d\n", string(idd[:]), tt)
	}
	log.Println("+++++++++++++++")
	hm.mutex.Lock()
	defer hm.mutex.Unlock()

	result := make([]*base.HttpInOut, 0)

	http_in_outs := hm.requests[id]
	for _, v := range http_in_outs {
		t, _ := v.Req.Base()
		if t == action {
			result = append(result, v)
		}
	}

	return result
}

func (hm *HttpMgr) Del(id [11]byte, action uint8) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()

	http_in_outs := hm.requests[id]
	for i, v := range http_in_outs {
		t, _ := v.Req.Base()
		if t == action {
			http_in_outs = append(http_in_outs[:i], http_in_outs[i+1:]...)
		}
	}
	if len(http_in_outs) == 0 {
		delete(hm.requests, id)
	}
}
