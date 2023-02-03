package controller

import (
	"truffle/mq"
	"truffle/ws/ddd/adapter"
	"truffle/ws/ddd/vo"
	"truffle/ws/service"
)

type MessageController struct {
	*service.MessageService
}

func InitMessageController(opt mq.MQOpt) {
	messageController = &MessageController{
		MessageService: service.NewMessageService(opt),
	}
}

func GetMessageController() *MessageController {
	return messageController
}

func (m *MessageController) GetMessage(req *vo.GetMsgRequest) *vo.GetMsgResponse {
	if req.Num <= 0 {
		req.Num = 1
	}
	return adapter.GetMsgResponseDto2Vo(m.MessageService.GetMessage(adapter.GetMsgRequestVo2Dto(req)))
}

func (m *MessageController) SendMsg(req *vo.SendMsgRequest) *vo.SendMsgResponse {
	m.MessageService.SendMsg(adapter.SendMsgRequestVo2Dto(req))
	return &vo.SendMsgResponse{}
}

var messageController *MessageController
