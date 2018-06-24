package reactor

import (
	"github.com/giskook/dtu/socket_server/protocol"
	"log"
)

//func (rc *Reactor) shunt() {
//	defer func() {
//		log.Println("ddddd")
//	}()
//
//	for {
//		select {
//		case <-rc.exit:
//			return
//		case socket_out := <-rc.socket_server.SocketOut:
//			_type, id := socket_out.Base()
//			hio := rc.hm.Get(id, rc.CT(_type))
//			for _, v := range hio {
//				if _type == protocol.PROTOCOL_2DSC_REGISTER {
//					v.Resp <- v.Req
//				}
//			}
//			rc.hm.Del(id, rc.CT(_type))
//		case http_inout := <-rc.http_srv.HttpInOut:
//			_, id := http_inout.Req.Base()
//			rc.hm.Put(id, http_inout)
//			rc.socket_server.SocketIn <- http_inout.Req
//		case http_cmd_del := <-rc.http_srv.HttpCmdDel:
//			rc.hm.Del(http_cmd_del.ID, http_cmd_del.Type)
//		}
//	}
//}
func (rc *Reactor) shunt() {
	defer func() {
		log.Println("ddddd")
	}()

	for {
		select {
		case <-rc.exit:
			return
		case socket_out := <-rc.socket_server.SocketOut:
			_type, id := socket_out.Base()
			log.Println(_type)
			log.Println(id)
		}
	}
}

func (rc *Reactor) CT(t uint8) uint8 {
	log.Printf(">>>>>> %d\n", t)
	switch t {
	case protocol.PROTOCOL_2DSC_REGISTER:
		return protocol.PROTOCOL_2DTU_REQ_REGISTER
	}

	return protocol.PROTOCOL_UNKNOWN
}
