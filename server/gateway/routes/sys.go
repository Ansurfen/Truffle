package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SysRoutes struct{}

func (router *SysRoutes) InstallSysApi(group *gin.RouterGroup) {
	sysRouter := group.Group("/sys")
	{
		sysRouter.GET("/metrics", Metrics(promhttp.Handler()))
		sysRouter.GET("/swagger/*any", Swagger())
	}
}

// @Summary 获取指标
// @Description 获取Prometheus统计的指标
// @Tags 系统相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /sys/metrics [get]
func Metrics(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// @Summary 获取API文档
// @Description 获取API文档
// @Tags 系统相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /sys/swagger/index.html [get]
func Swagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}
