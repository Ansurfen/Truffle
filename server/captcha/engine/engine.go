package engine

import (
	"github.com/gin-gonic/gin"
	. "truffle/breaker"
	. "truffle/captcha/proto"
	. "truffle/captcha/server"
	. "truffle/client"
	"truffle/etcd"
	. "truffle/etcd"
	truffle_i18n "truffle/i18n/proto"
	truffle_log "truffle/log/proto"
	"truffle/middleware"
	. "truffle/utils"
)

type CaptchaEngine struct {
	ec      *etcd.ESClient
	gtc     *GTClient
	httpSrv *gin.Engine
	gcs     map[string]*GClientConn
	ic      truffle_i18n.I18NClient
	lc      truffle_log.LogClient
}

func NewCaptchaEngine() *CaptchaEngine {
	opt := LoadOpt(ENV_DEVELOP, DefaultOpt{}, TracerOpt{}, EmailOpt{}, BreakerOpt{})
	if opt.Opt(ENV).(EnvOpt).Env == ENV_RELEASE {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := &CaptchaEngine{
		gtc:     NewGTClient(opt),
		ec:      NewESClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
		httpSrv: gin.Default(),
	}
	InitLoggerAdapter(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	return engine
}

func (engine *CaptchaEngine) Setup() *CaptchaEngine {
	engine.gcs = NewGClientConns([]string{"i18n", "log"}, WithTracer(engine.gtc.Tracer))
	engine.ic = engine.gcs["i18n"].NewI18NClient()
	engine.lc = engine.gcs["log"].NewLogClient()
	RegisterCaptchaServer(engine.gtc.Srv,
		NewCaptchaController(engine.ic, engine.lc, engine.gtc.Opt.Opt(EMAIL).(EmailOpt)))
	return engine
}

func (engine *CaptchaEngine) Mount() *CaptchaEngine {
	baseRouter := engine.httpSrv.Group("/captcha")
	baseRouter.Use(middleware.Cors())
	CaptchaRouter.InitImageApi(baseRouter)
	return engine
}

func (engine *CaptchaEngine) Run() {
	go engine.Shutdown()
	go engine.RunHttp()
	engine.gtc.Run()
}

func (engine *CaptchaEngine) RunHttp() {
	engine.httpSrv.Run(":" + engine.gtc.Opt.Opt(DEFAULT).(DefaultOpt).Service.HttpPort)
}

func (engine *CaptchaEngine) Shutdown() {
	engine.gtc.Shutdown()
	engine.ec.Close()
	for _, dial := range engine.gcs {
		defer dial.Close()
	}
	engine.gtc.Destory()
}
