package controller

import (
	"truffle/ws/ddd/adapter"
	"truffle/ws/ddd/vo"
	"truffle/ws/service"
)

type ChannelController struct {
	service *service.ChannelService
}

func NewChannelController() *ChannelController {
	return &ChannelController{
		service: service.NewChannelService(),
	}
}

func (con *ChannelController) GetChannels(req *vo.GetChannelRequest) *vo.GetChannelResponse {
	return adapter.GetChannelResponseDto2Vo(con.service.GetChannels(req.Path))
}
