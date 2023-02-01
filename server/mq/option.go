package mq

import (
	"truffle/utils"

	"github.com/Shopify/sarama"
)

const MQ = "mq"

type MQOpt struct {
	*sarama.Config
	Addr []string
}

func (opt MQOpt) Scheme() string {
	return MQ
}

func (opt MQOpt) Init(env string, c *utils.Conf) utils.IOpt {
	opt.Config = sarama.NewConfig()
	opt.Producer.RequiredAcks = sarama.WaitForAll
	opt.Producer.Partitioner = sarama.NewRandomPartitioner
	opt.Producer.Return.Successes = true
	switch env {
	case utils.ENV_DEVELOP:
		opt.Addr = c.GetStringSlice("kafka.develop.addr")
	case utils.ENV_RELEASE:
		opt.Addr = c.GetStringSlice("kafka.release.addr")
	}
	return opt
}
