package routes

import (
	"truffle/ws/hub"

	"github.com/gin-gonic/gin"
)

type WsServeRoutes struct{}

func (router *WsServeRoutes) InitWsApi(group *gin.RouterGroup) {
	wsRouter := group.Group("/ws")
	{
		wsRouter.GET("", WsServeHandle)
	}
}

func WsServeHandle(ctx *gin.Context) {
	hub.ServeWs(nil, ctx.Writer, ctx.Request)
}
