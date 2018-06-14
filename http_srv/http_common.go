package http_srv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/giskook/dtu/base"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type GeneralResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

var (
	GRS GeneralResponse = GeneralResponse{Code: base.ERR_NONE_CODE, Desc: base.ERR_NONE_DESC}
)

func EncodeErrResponse(w http.ResponseWriter, err *base.DtuError) {
	gr := &GeneralResponse{
		Code: err.Code,
		Desc: err.Desc(),
	}
	marshal_json(w, gr)
}

func RecordReq(r *http.Request) {
	v, e := httputil.DumpRequest(r, true)
	if e != nil {
		log.Println(e.Error())
		return
	}
	log.Println(string(v))
}

// MarshalJson 把对象以json格式放到response中
func marshal_json(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func unmarshal_json(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
