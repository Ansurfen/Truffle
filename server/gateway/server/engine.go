package server

import (
	"fmt"
	"os"
	"os/signal"
	. "truffle/client"
	. "truffle/etcd"
	truffle_user "truffle/user/proto"
	. "truffle/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

// ? proxy pattern

type GatewayEngine struct {
	httpSrv *gin.Engine
	opt     *BaseOpt
	ec      *EDClient
	gcs     map[string]*GClientConn
	uc      truffle_user.UserClient
}

func NewGatewayEngine() *GatewayEngine {
	opt := LoadOpt(ENV_DEVELOP, DefaultOpt{})
	if opt.Opt(ENV).(EnvOpt).Env == ENV_RELEASE {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := &GatewayEngine{
		opt:     opt,
		httpSrv: gin.Default(),
		ec:      NewEDClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
	}
	InitLoggerAdapter(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	resolver.Register(engine.ec)
	return engine
}

func (engine *GatewayEngine) Setup() *GatewayEngine {
	engine.gcs = NewGClientConns([]string{"user"})
	engine.uc = engine.gcs["user"].NewUserClient()
	return engine
}

func (engine *GatewayEngine) Mount() *GatewayEngine {
	engine.UserRouter()
	return engine
}

func (engine *GatewayEngine) Run() {
	go engine.Shutdown()
	zap.S().Info("Start to running...")
	if err := engine.httpSrv.Run(":" + engine.opt.Opt(DEFAULT).(DefaultOpt).Service.HttpPort); err != nil {
		fmt.Println(err)
	}
}

func (engine *GatewayEngine) Shutdown() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs
	for _, dial := range engine.gcs {
		defer dial.Close()
	}
	os.Exit(1)
}
