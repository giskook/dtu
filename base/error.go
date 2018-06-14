package base

import ()

type DtuError struct {
	Err      error
	Code     int
	Describe string
}

func (e *DtuError) Error() string {
	return e.Err.Error()
}

func (e *DtuError) Desc() string {
	return e.Describe
}

func NewErr(err error, code int, desc string) *DtuError {
	return &DtuError{
		Err:      err,
		Code:     code,
		Describe: desc,
	}
}

const (
	ERR_COMMON_NOT_CAPTURE_CODE int    = 999
	ERR_COMMON_NOT_CAPTURE_DESC string = "未捕获的错误"

	ERR_NONE_CODE int    = 0
	ERR_NONE_DESC string = "成功"

	ERR_HTTP_LACK_PARAMTERS_CODE int    = 1
	ERR_HTTP_LACK_PARAMTERS_DESC string = "缺少参数"

	ERR_HTTP_INNER_PANIC_CODE int    = 2
	ERR_HTTP_INNER_PANIC_DESC string = "内部错误"

	ERR_HTTP_TIMEOUT_CODE int    = 3
	ERR_HTTP_TIMEOUT_DESC string = "超时"

	ERR_DTU_OFFLINE_CODE int    = 4
	ERR_DTU_OFFLINE_DESC string = "DTU离线"

	ERR_DTU_PEER_CLOSE_CODE int    = 5
	ERR_DTU_PEER_CLOSE_DESC string = "DTU关闭了连接"
)

var (
	ERROR_HTTP_LACK_PARAMTERS *DtuError = NewErr(nil, ERR_HTTP_LACK_PARAMTERS_CODE, ERR_HTTP_LACK_PARAMTERS_DESC)
	ERROR_HTTP_INNER_PANIC    *DtuError = NewErr(nil, ERR_HTTP_INNER_PANIC_CODE, ERR_HTTP_INNER_PANIC_DESC)
	ERROR_HTTP_TIMEOUT        *DtuError = NewErr(nil, ERR_HTTP_TIMEOUT_CODE, ERR_HTTP_TIMEOUT_DESC)
	ERROR_NONE                *DtuError = NewErr(nil, ERR_NONE_CODE, ERR_NONE_DESC)
	ERROR_NOT_CAPTURE         *DtuError = NewErr(nil, ERR_COMMON_NOT_CAPTURE_CODE, ERR_COMMON_NOT_CAPTURE_DESC)
	ERROR_DTU_OFFLINE         *DtuError = NewErr(nil, ERR_DTU_OFFLINE_CODE, ERR_DTU_OFFLINE_DESC)
	ERROR_DTU_CLOSE_CONN      *DtuError = NewErr(nil, ERR_DTU_PEER_CLOSE_CODE, ERR_DTU_PEER_CLOSE_DESC)
)
