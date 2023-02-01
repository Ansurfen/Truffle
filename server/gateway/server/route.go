package server

import (
	"truffle/middleware"
	"truffle/utils"

	docs "truffle/gateway/docs"
	"truffle/gateway/routes"

	"github.com/gin-gonic/gin"
)

func (engine *GatewayEngine) UserRouter() *GatewayEngine {
	docs.SwaggerInfo.BasePath = "/"
	baseRouter := engine.httpSrv.Group("")
	baseRouter.Use(middleware.Cors())
	routes.GatewayRouter.InstallUserApi(baseRouter, engine.uc)
	routes.GatewayRouter.InstallSysApi(baseRouter)
	routes.GatewayRouter.InstallCaptchaApi(baseRouter)
	routes.GatewayRouter.InstallWsApi(baseRouter)
	engine.httpSrv.NoRoute(func(ctx *gin.Context) {
		utils.Proxy(ctx)
	})
	return engine
}
