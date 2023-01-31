package common

import (
	"encoding/json"

	"go.uber.org/zap"
)

type CommonResponse struct {
	Err  error `ddd:"*:err"` //TODO future tags
	Code int   `ddd:"*:code"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func newResult(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// ? clone pattern

func (res *Result) WithMsg(msg string) Result {
	return Result{
		Code: res.Code,
		Msg:  msg,
		Data: res.Data,
	}
}

func (res *Result) WithData(data any) Result {
	return Result{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

func (res *Result) WithJsonMarshal(data any) Result {
	raw, err := json.Marshal(data)
	if err != nil {
		zap.S().Warn(err)
		return JSON_MARSHAL_ERR
	}
	return Result{
		Code: res.Code,
		Msg:  res.Msg,
		Data: string(raw),
	}
}
