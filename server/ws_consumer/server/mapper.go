package server

import (
	"encoding/json"
	"truffle/db"
	"truffle/ws/ddd/po"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type WsConsumerMapper struct {
}

func NewWsConsumerMapper() *WsConsumerMapper {
	return &WsConsumerMapper{}
}

func (*WsConsumerMapper) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*WsConsumerMapper) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (*WsConsumerMapper) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for payload := range claim.Messages() {
		zap.S().Infof("Recv msg: %s %s\n", string(payload.Key), string(payload.Value))
		// switch mode
		// hungry
		var msg po.Message
		json.Unmarshal(payload.Value, &msg)
		db.GetDB().Model(po.Message{}).Create(msg)
		session.MarkMessage(payload, "")
		// lazy

	}
	return nil
}
