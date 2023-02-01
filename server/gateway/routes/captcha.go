package routes

import (
	"truffle/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type CaptchaRoutes struct{}

func (router *CaptchaRoutes) InstallCaptchaApi(group *gin.RouterGroup) {
	captchaRouter := group.Group("/captcha")
	{
		captchaRouter.GET("/image/get", GetImage())
		captchaRouter.GET("/image/get/:source", GetHistoryImage)
		captchaRouter.POST("/image/proof", Proof)
	}
}

// @Summary 获取图片验证码
// @Description 获取图片验证码
// @Tags 验证码相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /captcha/image/get [get]
func GetImage() gin.HandlerFunc {
	captchaImgCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "captchaImgCounter",
		Help: "count captcha image amount",
	})
	prometheus.MustRegister(captchaImgCounter)
	return func(ctx *gin.Context) {
		captchaImgCounter.Inc()
		utils.DoProxy(ctx.Writer, ctx.Request, utils.ProxyOpt{
			Addr: "http://localhost:9096",
		})
	}
}

// @Summary 获取历史图片验证码
// @Description 获取历史图片验证码
// @Tags 验证码相关
// @Accept json
// @Produce json
// @Param filename path string true "图片文件名"
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /captcha/image/get/{filename} [get]
func GetHistoryImage(ctx *gin.Context) {
	utils.DoProxy(ctx.Writer, ctx.Request, utils.ProxyOpt{
		Addr: "http://localhost:9096",
	})
}

// @Summary 图片验证码校验
// @Description 校验验证码是否正确
// @Tags 验证码相关
// @Accept json
// @Produce json
// @Param id formData string true "验证码的ID"
// @Param val formData string true "要校验的值"
// @Success 200 {string} string "{"status":"认证通过"}"
// @Failure 500 {string} string "{"status":"认证失败"}"
// @Router /captcha/image/proof [post]
func Proof(ctx *gin.Context) {
	utils.DoProxy(ctx.Writer, ctx.Request, utils.ProxyOpt{
		Addr: "http://localhost:9096",
	})
}
