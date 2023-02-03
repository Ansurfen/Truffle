package engine

import (
	"os"
	"os/signal"
	. "truffle/db"
	. "truffle/middleware"
	. "truffle/mq"
	. "truffle/utils"
	"truffle/ws/controller"
	"truffle/ws/ddd/po"
	"truffle/ws/hub"
	"truffle/ws/routes"

	"github.com/gin-gonic/gin"
)

type WsEngine struct {
	opt     *BaseOpt
	httpSrv *gin.Engine
}

func NewWsEngine() *WsEngine {
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, SQLOpt{}, NoSQLOpt{}, MQOpt{})
	if opt.Opt(ENV).(EnvOpt).Env == ENV_RELEASE {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := &WsEngine{
		opt:     opt,
		httpSrv: gin.Default(),
	}
	InitLogger(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	return engine
}

func (engine *WsEngine) Setup() *WsEngine {
	hub.InitHub()
	controller.InitWsController(hub.GetHub().GetAbstractHub())
	controller.InitMessageController(engine.opt.Opt(MQ).(MQOpt))
	return engine
}

func (engine *WsEngine) Mount() *WsEngine {
	opt := engine.opt
	InitDB(opt.Opt(SQL).(SQLOpt))
	InitRedis(opt.Opt(NOSQL).(NoSQLOpt))
	GetDB().AutoMigrate(&po.Channel{}, &po.CGroup{}, &po.Message{})
	baseRouter := engine.httpSrv.Group("")
	baseRouter.Use(Cors())
	routes.WsRouter.InitChannelApi(baseRouter)
	routes.WsRouter.InitMessageApi(baseRouter)
	routes.WsRouter.InitTopicApi(baseRouter)
	routes.WsRouter.InitWsApi(baseRouter)
	return engine
}

func (engine *WsEngine) Run() {
	go engine.Shutdown()
	go hub.GetHub().Run()
	engine.httpSrv.Run(":" + engine.opt.Opt(DEFAULT).(DefaultOpt).Service.HttpPort)
}

func (engine *WsEngine) Shutdown() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs
}
