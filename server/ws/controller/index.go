package controller

import (
	"truffle/ws/ddd/dao"
	"truffle/ws/ddd/po"
	"truffle/ws/ddd/pojo"
	"truffle/ws/ddd/vo"
	abstract "truffle/ws/interface"
)

type WsController struct {
	channelCon *ChannelController
	cgroupCon  *CGroupController
	topicCon   *TopicController
}

func InitWsController(hub abstract.AbstractHub) *WsController {
	abstract.InitGHub(hub)
	return &WsController{
		channelCon: NewChannelController(),
		cgroupCon:  NewCGroupController(),
		topicCon:   NewTopicController(),
	}
}

func GetWsController() *WsController {
	return wsController
}

func (ws *WsController) NewTopic(req *vo.NewTopicRequest) *vo.NewTopicResponse {
	res := ws.topicCon.NewTopic(req)
	ws.channelCon.service.ChannelMapper.Create(&dao.CreateChannelRequest{
		Name:  req.Name,
		Topic: req.Name,
		Path:  res.Path,
	})
	return &vo.NewTopicResponse{
		Path: abstract.GetHubWithAdapter().Encoding(res.Path),
	}
}

func (ws *WsController) JoinTopic(req *vo.JoinTopicRequest) *vo.JoinTopicResponse {
	cn := ws.channelCon.service.FindByPasscode(req.Passcode)
	ws.cgroupCon.CreateCGroup(req.User, "general", cn.Path)
	// 重复加入的情况判断
	// 没重复加入就返还cgroup, 虽然是加入channel,但是还是返回 cgroup
	// 还有channel属于是topic
	topic := po.Topic{
		Name: cn.Topic,
		Path: abstract.GetHubWithAdapter().Encoding(cn.Path),
	}
	return &vo.JoinTopicResponse{
		Topic: topic,
	}
}

func (ws *WsController) GetChannels(req *vo.GetChannelRequest) *vo.GetChannelResponse {
	req.Path = abstract.GetHubWithAdapter().Decoding(req.Path)
	return ws.channelCon.GetChannels(req)
}

func (ws *WsController) GetCGroup(req *vo.GetCGroupsRequset) *vo.GetCGroupsResponse {
	cgroups := ws.cgroupCon.CGroupMapper.FindCGroups(&dao.GetCGroupsRequset{
		User: req.User, Path: string(abstract.GetHubWithAdapter().Decoding(req.Path))})
	dict := make(map[string][]pojo.CGroupChan)
	for _, cgroupDao := range cgroups {
		dict[cgroupDao.Name] = append(dict[cgroupDao.Name], pojo.CGroupChan{
			Name: cgroupDao.Cname,
			Path: abstract.GetHubWithAdapter().Encoding(cgroupDao.Path),
		})
	}
	var dto []pojo.CGroupDTO
	for k, v := range dict {
		dto = append(dto, pojo.CGroupDTO{
			Title:    k,
			Channels: v,
		})
	}
	return &vo.GetCGroupsResponse{
		Cgroups: dto,
	}
}

var wsController *WsController
