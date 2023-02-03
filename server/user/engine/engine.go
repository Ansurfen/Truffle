package engine

import (
	. "truffle/client"
	. "truffle/db"
	. "truffle/etcd"
	truffle_i18n "truffle/i18n/proto"
	truffle_log "truffle/log/proto"
	. "truffle/user/proto"
	. "truffle/user/server"
	. "truffle/utils"

	"google.golang.org/grpc/resolver"
)

type UserEngine struct {
	ec  *EUClient
	gcs map[string]*GClientConn
	gtc *GTClient
	ic  truffle_i18n.I18NClient
	lc  truffle_log.LogClient
}

func NewUserEngine() *UserEngine {
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, TracerOpt{}, SQLOpt{}, NoSQLOpt{})
	engine := &UserEngine{
		gtc: NewGTClient(opt),
		ec:  NewEUClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
	}
	InitLogger(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	resolver.Register(engine.ec)
	return engine
}

// etcd and grpc
func (engine *UserEngine) Setup() *UserEngine {
	engine.gcs = NewGClientConns([]string{"i18n", "log"}, WithTracer(engine.gtc.Tracer))
	engine.ic = engine.gcs["i18n"].NewI18NClient()
	engine.lc = engine.gcs["log"].NewLogClient()
	RegisterUserServer(engine.gtc.Srv,
		NewUserController(engine.ic, engine.lc))
	opt := engine.gtc.Opt.Opt(DEFAULT).(DefaultOpt).Service
	engine.ec.Run(RegisterService(opt.Name, opt.Pod), opt.Addr, opt.Etcd.Expire)
	return engine
}

// middle mount
func (engine *UserEngine) Mount() *UserEngine {
	InitDB(engine.gtc.Opt.Opt(SQL).(SQLOpt))
	InitRedis(engine.gtc.Opt.Opt(NOSQL).(NoSQLOpt))
	GetDB().AutoMigrate(User{})
	return engine
}

func (engine *UserEngine) Run() {
	go engine.Shutdown()
	engine.gtc.Run()
}

func (engine *UserEngine) Shutdown() {
	engine.gtc.Shutdown()
	engine.ec.Close()
	for _, dial := range engine.gcs {
		defer dial.Close()
	}
	engine.gtc.Destory()
}
