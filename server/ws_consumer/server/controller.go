package server

import (
	"truffle/client"
	"truffle/mq"
	truffle_ws "truffle/ws/proto"
)

type WsConsumerController struct {
	wc truffle_ws.WsClient
	*WsConsumerMapper
	consumer *mq.MQCosumerGroup
}

func NewWsConsumerController(opt mq.MQOpt, wc truffle_ws.WsClient) *WsConsumerController {
	return &WsConsumerController{
		wc:               wc,
		consumer:         mq.NewMQConsumerGroup(opt),
		WsConsumerMapper: NewWsConsumerMapper(),
	}
}

func (con *WsConsumerController) WriteMsgByRPC() {
	client.WriteMsg(con.wc, func(res *truffle_ws.WriteMsgResponse) error {
		return nil
	})
	return
}

func (con *WsConsumerController) WriteMsgByMQ() {
	con.consumer.Consume([]string{"truffle_send_msg"}, con.WsConsumerMapper)
}

func (con *WsConsumerController) Close() {
	con.consumer.Close()
}