package vo

import "truffle/ws/ddd/pojo"

type GetCGroupsResponse struct {
	Cgroups []pojo.CGroupDTO
}

type GetCGroupsRequset struct {
	Path string `json:"path" form:"path"`
	User string `json:"user" form:"user"`
}
