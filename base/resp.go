package base

import (
	"ex/constant"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SucResp(data any) *Resp {
	return &Resp{Code: constant.Success, Data: data}
}

func UnknownResp() *Resp {
	return &Resp{Code: constant.Unknown}
}

func ErrorResp(code int) *Resp {
	return &Resp{Code: code}
}
