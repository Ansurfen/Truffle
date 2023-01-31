package server

import (
	"math/rand"
	"net/http"
	"path"
	"time"
	"truffle/breaker"
	"truffle/common"
	"truffle/utils"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
)

var (
	ids           map[string]bool
	CaptchaRouter = new(CaptchaImageRoutes)
)

type CaptchaImageRoutes struct{}

func (router *CaptchaImageRoutes) InitImageApi(group *gin.RouterGroup) {
	imageRouter := group.Group("/image")
	imageRouter.Use(sentinelPlugin.SentinelMiddleware())
	{
		breaker.LoadRule("GET:/captcha/image/get", 100, uint32(time.Second))
		_, b := sentinel.Entry("captcha", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
		} else {
			imageRouter.GET("/get", GetImage)
		}
		imageRouter.GET("/get/:source", GetHistoryImage)
		imageRouter.POST("/proof", Proof)
	}
}

func GetImage(ctx *gin.Context) {
	id := utils.GetCaptcha()
	if ids == nil {
		ids = make(map[string]bool)
	}
	ids[id] = false
	utils.SetImage(ctx.Writer, ctx.Request, id)
}

func GetHistoryImage(ctx *gin.Context) {
	_, file := path.Split(ctx.Request.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if _, ok := ids[id]; ok {
		utils.SetFile(ctx.Writer, ctx.Request)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Image does not exist or has expired"})
	}
}

func Proof(ctx *gin.Context) {
	common.GinWrap(func(req *ProofRequest) common.Result {
		delete(ids, req.Id)
		if utils.CaptchaProof(req.Id, req.Val) {
			return common.SUCCESS.WithMsg("pass")
		}
		return common.ERROR.WithMsg("unpass")
	})
}
