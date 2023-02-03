package vo

import (
	"truffle/common"
	"truffle/ws/ddd/po"
)

// @SendMsgRequestVo2Dto req:v2d message.go
type SendMsgRequest struct {
	Path string     `json:"path" ddd:"v2d:path"`
	Msg  po.Message `json:"msg" ddd:"v2d:msg"`
}

type SendMsgResponse struct {
	Data string
}

// @GetMsgRequestVo2Dto req:v2d message.go
type GetMsgRequest struct {
	Path string `json:"path" ddd:"v2d:path"`
	Num  int64  `json:"id" ddd:"v2d:num"`
}

// @GetMsgResponseDto2Vo ret:d2v message.go
type GetMsgResponse struct {
	Meta common.CommonResponse
	Msgs []po.Message `json:"msgs"`
	Id   int          `json:"id"`
}

type DelMsgRequest struct {
	Path string
}

type DelMsgResponse struct {
	Meta common.CommonResponse
}
