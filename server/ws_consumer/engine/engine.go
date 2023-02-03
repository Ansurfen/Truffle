package engine

import (
	"os"
	"os/signal"
	. "truffle/client"
	. "truffle/db"
	. "truffle/etcd"
	. "truffle/mq"
	. "truffle/utils"
	. "truffle/ws_consumer/server"

	"google.golang.org/grpc/resolver"
)

type WsConsumerEngine struct {
	ec       *EUClient
	opt      *BaseOpt
	gtc      *GTClient
	gcs      map[string]*GClientConn
	consumer *WsConsumerController
}

func NewWsConsumerEngine() *WsConsumerEngine {
	opt := LoadOpt(ENV_DEVELOP, DefaultOpt{}, TracerOpt{}, SQLOpt{}, MQOpt{})
	engine := &WsConsumerEngine{
		opt: opt,
		gtc: NewGTClient(opt),
		ec:  NewEUClient(opt.Opt(DEFAULT).(DefaultOpt).Service.Etcd.Addr),
	}
	resolver.Register(engine.ec)
	return engine
}

func (engine *WsConsumerEngine) Setup() *WsConsumerEngine {
	engine.gcs = NewGClientConns([]string{"ws"}, WithTracer(engine.gtc.Tracer))
	engine.consumer = NewWsConsumerController(engine.gtc.Opt.Opt(MQ).(MQOpt),
		engine.gcs["ws"].NewWsClient())
	return engine
}

func (engine *WsConsumerEngine) Mount() *WsConsumerEngine {
	InitDB(engine.opt.Opt(SQL).(SQLOpt))
	return engine
}

func (engine *WsConsumerEngine) Run() {
	go engine.Shutdown()
	go engine.consumer.WriteMsgByMQ()
	engine.gtc.Run()
}

func (engine *WsConsumerEngine) Shutdown() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs
	engine.ec.Close()
	for _, dial := range engine.gcs {
		defer dial.Close()
	}
	engine.consumer.Close()
	engine.gtc.Destory()
}
