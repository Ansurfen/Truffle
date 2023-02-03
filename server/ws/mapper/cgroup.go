package mapper

import (
	"truffle/db"
	"truffle/utils"
	"truffle/ws/ddd/dao"
	"truffle/ws/ddd/po"
)

type CGroupMapper struct {
	Node *utils.SnowFlake
}

func NewCGroupMapper() *CGroupMapper {
	return &CGroupMapper{
		Node: utils.NewSnowFlake(1),
	}
}

func (mapper *CGroupMapper) CreateCGroup(user, name, path string) {
	db.GetDB().Model(&po.CGroup{}).Create(&po.CGroup{
		Id:   mapper.Node.Generate().Int64(),
		User: user,
		Name: name,
		Path: path,
	})
}

func (mapper *CGroupMapper) FindCGroup(user, path string) []po.CGroup {
	var cgroups []po.CGroup
	db.GetDB().Model(&po.CGroup{}).Where("user = ? AND path = ?", user, path).Find(&cgroups)
	return cgroups
}

func (mapper *CGroupMapper) FindCGroups(req *dao.GetCGroupsRequset) []dao.GetCGroupResponse {
	var res []dao.GetCGroupResponse
	db.GetDB().Raw(`select cgroup.name, 
	cgroup.path, channel.name as cname from cgroup join channel 
	on cgroup.path = channel.path and cgroup.user = ? and cgroup.path like ?`, req.User, req.Path+"%").Scan(&res)
	return res
}
