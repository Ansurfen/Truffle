package engine

import (
	. "truffle/etcd"
	truffle_log "truffle/log/proto"

	. "truffle/client"
	. "truffle/log/server"
	. "truffle/mq"
	. "truffle/utils"
)

type LogEngine struct {
	gtc *GTClient
	ec  *ESClient
	mq  *MQProducer
}

func NewLogEngine() *LogEngine {
	opt := LoadOpt(ENV_DEVELOP, EnvOpt{}, DefaultOpt{}, TracerOpt{}, MQOpt{})
	engine := &LogEngine{
		gtc: NewGTClient(opt),
		ec:  NewESClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
	}
	InitLoggerAdapter(opt.Opt(DEFAULT).(DefaultOpt).Logger)
	return engine
}

func (engine *LogEngine) Setup() *LogEngine {
	opt := engine.gtc.Opt.Opt(DEFAULT).(DefaultOpt).Service
	engine.ec.Run(RegisterService(opt.Name, opt.Pod), opt.Addr, opt.Etcd.Expire)
	return engine
}

func (engine *LogEngine) Mount() *LogEngine {
	logController := NewLogController()
	if engine.gtc.Opt.Opt(ENV).(EnvOpt).UseMQ {
		opt := engine.gtc.Opt.Opt(MQ).(MQOpt)
		logController.Mq = NewMQProducer(opt.Addr, &opt)
		engine.mq = logController.Mq
	}
	truffle_log.RegisterLogServer(engine.gtc.Srv, logController)
	return engine
}

func (engine *LogEngine) Run() {
	go engine.Shutdown()
	engine.gtc.Run()
}

func (engine *LogEngine) Shutdown() {
	engine.gtc.Shutdown()
	engine.ec.Close()
	if engine.mq != nil {
		engine.mq.Close()
	}
	engine.gtc.Destory()
}
