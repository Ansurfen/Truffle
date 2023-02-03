package dto

import (
	"truffle/common"
	"truffle/ws/ddd/po"
)

// @GetMsgRequestDto2Dao req:d2d message.go
// @GetMsgRequestVo2Dto ret:dto message.go
type GetMsgRequest struct {
	Path string `ddd:"d2d:path"`
	Num  int64  `ddd:"d2d:base"`
}

// @GetMsgResponseDto2Vo req:d2v message.go
// @GetMsgResponseDao2Dto ret:d2d message.go
type GetMsgResponse struct {
	Meta common.CommonResponse `ddd:"*:meta"`
	Num  int64
	Msgs []po.Message
}

// @SendMsgRequestVo2Dto ret:v2d message.go
type SendMsgRequest struct {
	Path string
	Msg  po.Message
}

type SendMsgResponse struct {
	Meta common.CommonResponse
}
