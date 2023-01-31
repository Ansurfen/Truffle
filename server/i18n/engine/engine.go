package engine

import (
	. "truffle/client"
	. "truffle/etcd"
	. "truffle/i18n/proto"
	. "truffle/i18n/server"
	. "truffle/i18n/store"
	. "truffle/utils"
)

type I18NEngine struct {
	gtc *GTClient
	ec  *ESClient
}

func NewI18NEngine() *I18NEngine {
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, TracerOpt{}, LangOpt{})
	engine := &I18NEngine{
		gtc: NewGTClient(opt),
		ec:  NewESClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
	}
	InitLoggerAdapter(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	Lopt = opt.Opt(LANG).(LangOpt)
	return engine
}

func (engine *I18NEngine) Setup() *I18NEngine {
	RegisterI18NServer(engine.gtc.Srv, NewI18NController())
	opt := engine.gtc.Opt.Opt(DEFAULT).(DefaultOpt).Service
	engine.ec.Run(RegisterService(opt.Name, opt.Pod), opt.Addr, opt.Etcd.Expire)
	return engine
}

func (engine *I18NEngine) Mount() *I18NEngine {
	return engine
}

func (engine *I18NEngine) Run() {
	go engine.Shutdown()
	engine.gtc.Run()
}

func (engine *I18NEngine) Shutdown() {
	engine.gtc.Shutdown()
	engine.ec.Close()
	engine.gtc.Destory()
}
