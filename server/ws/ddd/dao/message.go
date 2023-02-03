package dao

import (
	"truffle/common"
	"truffle/ws/ddd/po"
)

// @GetMsgRequestDto2Dao ret:dao message.go
type GetMsgRequest struct {
	Path     string
	Base     int64
	Limit    int
	PageSize int
}

// @GetMsgResponseDao2Dto req:d2d message.go
type GetMsgResponse struct {
	Meta common.CommonResponse `ddd:"*:meta"`
	Msgs []po.Message          `ddd:"d2d:msgs"`
}
