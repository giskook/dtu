package http_srv

import (
	"context"
	"github.com/giskook/dtu/base"
	"net/http"
	"time"
)

func (h *HttpSrv) handler_web_user_add(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w, base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	r.ParseForm()
	defer r.Body.Close()

	id := r.Form.Get(WEB_USER_MODIFY_PARA_ID)
	alias := r.Form.Get(WEB_USER_MODIFY_PARA_ALIAS)
	user_type := r.Form.Get(WEB_USER_ADD_PARA_USER_TYPE)
	if id == "" ||
		alias == "" ||
		user_type == "" {
		EncodeErrResponse(w, base.ERROR_HTTP_LACK_PARAMTERS)
	}

	ctx := r.Context()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	user_add_result := make(chan *base.DBResult)
	go h.db.UserAdd(id, alias, id, user_type, user_add_result)
	select {
	case <-ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
	case result := <-user_add_result:
		if result.Err != nil {
			EncodeErrResponse(w, result.Err.(*base.LorawanError))
			return
		}
		EncodeErrResponse(w, base.ERROR_NONE)
	}
}
