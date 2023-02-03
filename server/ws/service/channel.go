package service

import (
	"truffle/ws/ddd/dto"
	"truffle/ws/mapper"
)

type ChannelService struct {
	*mapper.ChannelMapper
}

func NewChannelService() *ChannelService {
	return &ChannelService{}
}

func (service *ChannelService) GetChannels(path string) *dto.GetChannelsResponse {
	return &dto.GetChannelsResponse{
		Channels: service.ChannelMapper.FindByPath(path),
	}
}