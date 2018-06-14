package http_srv

import (
	"context"
	"github.com/giskook/dtu/base"
	"net/http"
	"time"
)

const (
	HTTP_DTU_ID string = "plc_id"
)

func (h *HttpSrv) handler_dtu_restart(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w, base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	id := r.Form.Get(HTTP_DTU_ID)
	var dtu_id [11]byte
	copy(dtu_id[:], []byte(id))
	if id == "" {
		EncodeErrResponse(w, base.ERROR_HTTP_LACK_PARAMTERS)
	}

	ctx := r.Context()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	result := make(chan base.Proto)
	h.HttpInOut <- &base.HttpInOut{
		Req: &base.Restart{
			Type: base.PROTOCOL_2DTU_REQ_REGISTER,
			ID:   dtu_id,
		},
		Resp: result,
	}

	select {
	case <-ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
		h.HttpCmdDel <- &base.InnerCmdDel{
			Type: base.PROTOCOL_2DTU_REQ_REGISTER,
			ID:   dtu_id,
		}
	case <-result:
		EncodeErrResponse(w, base.ERROR_NONE)
	}
}
