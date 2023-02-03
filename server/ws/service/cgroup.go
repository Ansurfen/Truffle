package service

import "truffle/ws/mapper"

type CGroupService struct {
	*mapper.CGroupMapper
}

func NewCGroupService() *CGroupService {
	return &CGroupService{
		CGroupMapper: mapper.NewCGroupMapper(),
	}
}

func (service *CGroupService) CreateCGroup(user, name, path string) {
	service.CGroupMapper.CreateCGroup(user, name, path)
}
