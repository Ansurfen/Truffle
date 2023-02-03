package adapter

import (
	"truffle/ws/ddd/dto"
	"truffle/ws/ddd/vo"
)

func GetChannelResponseDto2Vo(req *dto.GetChannelsResponse) *vo.GetChannelResponse {
	 return &vo.GetChannelResponse{
		Channels:	req.Channels,
	}
}

