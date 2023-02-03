package routes

import (
	"truffle/common"
	"truffle/ws/controller"
	"truffle/ws/ddd/vo"

	"github.com/gin-gonic/gin"
)

type WsTopicRoutes struct{}

func (router *WsTopicRoutes) InitTopicApi(group *gin.RouterGroup) {
	messageRouter := group.Group("/topic")
	{
		messageRouter.POST("/new", NewTopicHandle)
		messageRouter.POST("/join", JoinTopicHandle)
	}
}

func JoinTopicHandle(ctx *gin.Context) {
	common.GinWrap(func(req vo.JoinTopicRequest) common.Result {
		res := controller.GetWsController().JoinTopic(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}

func NewTopicHandle(ctx *gin.Context) {
	common.GinWrap(func(req vo.NewTopicRequest) common.Result {
		res := controller.GetWsController().NewTopic(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}
