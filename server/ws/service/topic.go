package service

import (
	"truffle/ws/ddd/dto"
	abstract "truffle/ws/interface"
	"truffle/ws/mapper"
)

type TopicService struct {
	mapper *mapper.TopicMapper
}

func NewTopicService() *TopicService {
	return &TopicService{
		mapper: mapper.NewTopicMapper(),
	}
}

func (service *TopicService) NewTopic(req *dto.NewTopicRequest) *dto.NewTopicResponse {
	path := abstract.GetHubWithAdapter().GenPath(req.Name)
	return &dto.NewTopicResponse{
		Path: path,
	}
}

func (service *TopicService) JoinTopic(req *dto.JoinTopicRequest) *dto.JoinTopicResponse {
	return &dto.JoinTopicResponse{}
}
