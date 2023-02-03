package mapper

import (
	"truffle/db"
	"truffle/utils"
	"truffle/ws/ddd/dao"
	"truffle/ws/ddd/po"
)

type ChannelMapper struct {
	Node *utils.SnowFlake
}

func NewChannelMapper() *ChannelMapper {
	return &ChannelMapper{
		Node: utils.NewSnowFlake(1),
	}
}

func (mapper *ChannelMapper) Create(req *dao.CreateChannelRequest) {
	db.GetDB().Model(&po.Channel{}).Create(&po.Channel{
		Id:       mapper.Node.Generate().Int64(),
		Topic:    req.Topic,
		Path:     req.Path,
		Name:     req.Name,
		Passcode: utils.MD5(req.Topic + req.Path + req.Name),
	})
}

func (mapper *ChannelMapper) FindByPath(path string) []po.Channel {
	var channels []po.Channel
	db.GetDB().Model(&po.Channel{}).Where("path = ?", path).Find(&channels)
	return channels
}

func (mapper *ChannelMapper) FindByPasscode(code string) po.Channel {
	var channel po.Channel
	db.GetDB().Model(&po.Channel{}).Where("passcode = ?", code).First(&channel)
	return channel
}
