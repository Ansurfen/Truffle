package routes

import (
	"truffle/common"
	"truffle/ws/controller"
	"truffle/ws/ddd/vo"

	"github.com/gin-gonic/gin"
)

type WsChannelRoutes struct{}

func (router *WsChannelRoutes) InitChannelApi(group *gin.RouterGroup) {
	channelRouter := group.Group("/channel")
	{
		channelRouter.POST("/get", GetChannels)
		channelRouter.POST("/group/get", GetCGroups)
	}
}

func GetChannels(ctx *gin.Context) {
	common.GinWrap(func(req vo.GetChannelRequest) common.Result {
		res := controller.GetWsController().GetChannels(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}

func GetCGroups(ctx *gin.Context) {
	common.GinWrap(func(req vo.GetCGroupsRequset) common.Result {
		res := controller.GetWsController().GetCGroup(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}
