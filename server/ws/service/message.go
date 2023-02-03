package service

import (
	"encoding/json"
	"truffle/common"
	"truffle/mq"
	"truffle/ws/ddd/adapter"
	"truffle/ws/ddd/dto"
	"truffle/ws/mapper"
)

type MessageService struct {
	*mapper.MessageMapper
	*mq.MQProducer
}

func NewMessageService(opt mq.MQOpt) *MessageService {
	return &MessageService{
		MessageMapper: mapper.NewMessageMapper(),
		MQProducer:    mq.NewMQProducer(opt.Addr, &opt),
	}
}

func (service *MessageService) SendMsg(req *dto.SendMsgRequest) *dto.SendMsgResponse {
	msg := req.Msg.AdapterDB(service.MessageMapper.Node.Generate().Int64())
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return &dto.SendMsgResponse{
			Meta: common.CommonResponse{
				Err:  err,
				Code: common.JSON_MARSHAL_ERR_CODE,
			},
		}
	}
	id, err := service.MessageMapper.CreateMsgCache(req.Path, string(msgStr))
	if err != nil {
		// write log
		// or write db
		service.MessageMapper.CreateMessage(msg)
		return &dto.SendMsgResponse{
			Meta: common.CommonResponse{
				Err:  err,
				Code: common.ERROR_BUT_NO_ISSUCE_CODE,
			},
		}
	}
	service.MQProducer.SendMessage(mq.NewMQMessage("truffle_send_msg", "tf_id_"+id, string(msgStr)))
	return &dto.SendMsgResponse{
		Meta: common.CommonResponse{
			Code: common.SUCCESS_CODE,
		},
	}
}

func (service *MessageService) DeleteMessage() {

}

func (service *MessageService) GetMessage(req *dto.GetMsgRequest) *dto.GetMsgResponse {
	dao := adapter.GetMsgRequestDto2Dao(req)
	dao.PageSize = 15
	dao.Limit = 15
	ret := adapter.GetMsgResponseDao2Dto(service.MessageMapper.GetMessage(dao))
	for idx, msg := range ret.Msgs {
		ret.Msgs[idx] = msg.AdapterIMsg()
	}
	if len(ret.Msgs) > 0 {
		ret.Num++
	}
	return ret
}
