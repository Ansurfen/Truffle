package mq

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type MQProducer struct {
	sarama.SyncProducer
	opt *MQOpt
}

func NewMQProducer(addrs []string, opt *MQOpt) *MQProducer {
	cli, err := sarama.NewSyncProducer(addrs, opt.Config)
	if err != nil {
		zap.S().Fatalf("Producer close, err: %v", err)
		return nil
	}
	return &MQProducer{
		SyncProducer: cli,
		opt:          opt,
	}
}

func (mq *MQProducer) SendMsg(msg *sarama.ProducerMessage) {
	pid, offset, err := mq.SendMessage(msg)
	if err != nil {
		zap.S().Warnf("Fail to send message, err: %v", err)
		return
	}
	zap.S().Debugf("pid: %v offset: %v", pid, offset)
}
