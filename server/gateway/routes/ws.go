package routes

import "github.com/gin-gonic/gin"

type WsRoutes struct{}

func (router *WsRoutes) InstallWsApi(group *gin.RouterGroup) {
	wsRouter := group.Group("")
	{
		wsRouter.GET("/ws")
		msgRouter := wsRouter.Group("/message")
		{
			msgRouter.POST("/get", GetMsgHandle)
			msgRouter.POST("/send", SendMsgHandle)
		}
		topicRouter := wsRouter.Group("/topic")
		{
			topicRouter.POST("/new", NewTopicHandle)
			topicRouter.POST("/join", JoinTopicHandle)
		}
		channelRouter := wsRouter.Group("/channel")
		{
			channelRouter.POST("/get", GetChannels)
			channelRouter.POST("/group/get", GetCGroups)
		}
	}
}

// @Summary 新增话题
// @Description 新增话题
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /topic/new [post]
func NewTopicHandle(ctx *gin.Context) {}

// @Summary 加入话题
// @Description 加入话题
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /topic/join [post]
func JoinTopicHandle(ctx *gin.Context) {}

// @Summary 获取消息
// @Description 获取消息
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /message/get [post]
func GetMsgHandle(ctx *gin.Context) {}

// @Summary 发送消息
// @Description 发送消息
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /message/send [post]
func SendMsgHandle(ctx *gin.Context) {}

// @Summary 获取频道
// @Description 获取频道
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /channel/get [post]
func GetChannels(ctx *gin.Context) {}

// @Summary 获取频道组
// @Description 获取频道组
// @Tags 即时通讯相关
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Failure 500 {string} Fail
// @Router /channel/group/get [post]
func GetCGroups(ctx *gin.Context) {}
