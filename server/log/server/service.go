package service

import (
	"go.uber.org/zap"
	. "truffle/log/proto"
	. "truffle/mq"
)

type LogService struct{}

func (s *LogService) PrintLogger(level, msg string) *LogResponse {
	zap.S().Infof("Level: %s Msg: %s", level, msg)
	return &LogResponse{Ok: true}
}

func (s *LogService) PushMQ(topic, level, msg string, mq *MQProducer) *LogResponse {
	if mq == nil {
		zap.S().Warn("Fail to init mq")
		return &LogResponse{Ok: false}
	}
	mq.SendMsg(NewMQMessage(topic, level, msg))
	return &LogResponse{Ok: true}
}
