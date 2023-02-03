package vo

import "truffle/ws/ddd/po"

type GetChannelRequest struct {
	Path string `json:"path" form:"path"`
}

// @GetChannelResponseDto2Vo ret:d2v channel.go
type GetChannelResponse struct {
	Channels []po.Channel
}
