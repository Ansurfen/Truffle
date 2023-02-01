package mq

import (
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
