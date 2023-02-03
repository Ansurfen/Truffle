package adapter

import (
	"truffle/ws/ddd/vo"
	"truffle/ws/ddd/dto"
)

func JoinTopicRequestVo2Dto(req *vo.JoinTopicRequest) *dto.JoinTopicRequest {
	 return &dto.JoinTopicRequest{}
}

func NewTopicRequestVo2Dto(req *vo.NewTopicRequest) *dto.NewTopicRequest {
	 return &dto.NewTopicRequest{}
}

