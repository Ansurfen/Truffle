package controller

import (
	"truffle/ws/ddd/adapter"
	"truffle/ws/ddd/vo"
	"truffle/ws/service"
)

type TopicController struct {
	service *service.TopicService
}

func NewTopicController() *TopicController {
	return &TopicController{
		service: service.NewTopicService(),
	}
}

func (con *TopicController) NewTopic(req *vo.NewTopicRequest) *vo.NewTopicResponse {
	con.service.NewTopic(adapter.NewTopicRequestVo2Dto(req))
	return &vo.NewTopicResponse{}
}

func (con *TopicController) JoinTopic(req *vo.JoinTopicRequest) *vo.JoinTopicResponse {
	con.service.JoinTopic(adapter.JoinTopicRequestVo2Dto(req))
	return &vo.JoinTopicResponse{}
}
