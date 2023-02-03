package controller

import (
	"truffle/ws/ddd/vo"
	"truffle/ws/service"
)

type CGroupController struct {
	*service.CGroupService
}

func NewCGroupController() *CGroupController {
	return &CGroupController{
		CGroupService: service.NewCGroupService(),
	}
}

func (con *CGroupController) GetCGroup(req *vo.GetCGroupsRequset) *vo.GetCGroupsResponse {
	return &vo.GetCGroupsResponse{}
}

func (con *CGroupController) CreateCGroup(user, name, path string) {
	con.CGroupService.CreateCGroup(user, name, path)
}
