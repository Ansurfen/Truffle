package dto

import "truffle/ws/ddd/po"

type GetChannelsRequest struct {
	Path string
}

// @GetChannelResponseDto2Vo req:d2v channel.go
type GetChannelsResponse struct {
	Channels []po.Channel `ddd:"d2v:channels"`
}
