package mq

import "github.com/Shopify/sarama"

func NewMQMessage(topic, k, v string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(k),
		Value: sarama.StringEncoder(v),
	}
}
