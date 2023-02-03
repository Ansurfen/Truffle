package mq

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type MQConsumer struct {
	sarama.Consumer
	wg sync.WaitGroup
}

func NewMQConsumer(addrs []string) *MQConsumer {
	cli, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		zap.S().Fatalf("Consumer connnet err: %v", err)
		return nil
	}
	return &MQConsumer{
		Consumer: cli,
	}
}

func (mq *MQConsumer) GetPartitions(topic string) []int32 {
	partitions, err := mq.Consumer.Partitions(topic)
	if err != nil {
		zap.S().Warnf("Fail to get paritions, err: %v", err)
		return []int32{}
	}
	return partitions
}

func (mq *MQConsumer) Comsume(topic string, f func(*sarama.ConsumerMessage)) {
	partitions, err := mq.Consumer.Partitions(topic)
	if err != nil {
		zap.S().Warnf("Fail to get paritions, err: %v", err)
		return
	}
	for _, partition := range partitions {
		partitionConsumer, err := mq.Consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			zap.S().Warnf("partitionConsumer err: %v", err)
			continue
		}
		mq.wg.Add(1)
		go func() {
			for msg := range partitionConsumer.Messages() {
				f(msg)
			}
			mq.wg.Done()
		}()
	}
	mq.wg.Wait()
}

type MQCosumerGroup struct {
	conf MQOpt
	sarama.ConsumerGroup
}

func NewMQConsumerGroup(opt MQOpt) *MQCosumerGroup {
	opt.Config.Consumer.Return.Errors = false
	opt.Config.Version = sarama.V0_10_2_0
	opt.Config.Consumer.Offsets.Initial = sarama.OffsetOldest
	group, err := sarama.NewConsumerGroup(opt.Addr, "", opt.Config)
	if err != nil {
		zap.S().Fatal(err)
	}
	consumer := &MQCosumerGroup{
		conf:          opt,
		ConsumerGroup: group,
	}
	return consumer
}

func (consumer *MQCosumerGroup) TrackErr() {
	for err := range consumer.ConsumerGroup.Errors() {
		zap.S().Warnf("err: %v", err)
	}
}

func (consumer *MQCosumerGroup) Consume(topics []string, handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := consumer.ConsumerGroup.Consume(ctx, topics, handler)
		if err != nil {
			zap.S().Warn(err)
		}
	}
}
