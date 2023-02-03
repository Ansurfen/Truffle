package routes

import (
	"truffle/common"
	"truffle/ws/controller"
	"truffle/ws/ddd/vo"

	"github.com/gin-gonic/gin"
)

type WsMessageRoutes struct{}

func (router *WsMessageRoutes) InitMessageApi(group *gin.RouterGroup) {
	messageRouter := group.Group("/message")
	{
		messageRouter.POST("/get", GetMsgHandle)
		messageRouter.POST("/send", SendMsgHandle)
	}
}

func GetMsgHandle(ctx *gin.Context) {
	common.GinWrap(func(req vo.GetMsgRequest) common.Result {
		res := controller.GetMessageController().GetMessage(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}

func SendMsgHandle(ctx *gin.Context) {
	common.GinWrap(func(req vo.SendMsgRequest) common.Result {
		res := controller.GetMessageController().SendMsg(&req)
		return common.SUCCESS.WithJsonMarshal(res)
	})(ctx)
}
