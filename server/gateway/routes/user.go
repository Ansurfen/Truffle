package routes

import (
	"truffle/client"
	"truffle/common"
	"truffle/gateway/models"
	truffle_user "truffle/user/proto"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type UserRoutes struct{}

func (router *UserRoutes) InstallUserApi(group *gin.RouterGroup, uc truffle_user.UserClient) {
	userRouter := group.Group("/user")
	{
		userRouter.POST("/login", Login(uc))
		userRouter.POST("/register", Register(uc))
	}
}

// @Summary 用户登录
// @Description 通过RPC对话验证码服务。判断验证码是否复合要求在决定是否放行，网关不鉴别JWT留给用户服务实现，主要怕别人伪造接口，登录成功会返回ws服务器的地址
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param key formData string true "邮箱/手机号/用户名"
// @Param pwd formData string true "登录密码"
// @Success 200 {string} string "{"status":"登录成功"}"
// @Failure 500 {string} string "{"status":"登录失败"}"
// @Router /user/login [post]
func Login(uc truffle_user.UserClient) gin.HandlerFunc {
	loginCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "loginCounter",
		Help: "count login amount",
	})
	prometheus.MustRegister(loginCounter)
	return func(ctx *gin.Context) {
		loginCounter.Inc()
		common.GinWrap(func(req models.LoginRequest) common.Result {
			res := client.Login(uc, req.Key, req.Pwd)
			if !res.Ok {
				return common.ERROR.WithData(res)
			}
			return common.SUCCESS.WithData(res)
		})(ctx)
	}
}

// @Summary 用户注册
// @Description 注册新用户
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param name formData string true "用户名"
// @Param pwd formData string true "用户密码"
// @Param key formData string true "邮箱"
// @Success 200 {string} string "{"status":"注册成功"}"
// @Failure 500 {string} string "{"status":"注册失败"}"
// @Router /user/register [post]
func Register(uc truffle_user.UserClient) gin.HandlerFunc {
	registerCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "registerCounter",
		Help: "count register amount",
	})
	prometheus.MustRegister(registerCounter)
	return func(ctx *gin.Context) {
		registerCounter.Inc()
		common.GinWrap(func(req models.RegisterRequest) common.Result {
			res := client.Register(uc, "zh_cn", req.Name, req.Pwd, req.Key)
			if !res.Ok {
				return common.ERROR.WithData(res)
			}
			return common.SUCCESS.WithData(res)
		})(ctx)
	}
}
