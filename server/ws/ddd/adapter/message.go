package adapter

import (
	"truffle/ws/ddd/dao"
	"truffle/ws/ddd/dto"
	"truffle/ws/ddd/vo"
)

func SendMsgRequestVo2Dto(req *vo.SendMsgRequest) *dto.SendMsgRequest {
	return &dto.SendMsgRequest{
		Path: req.Path,
		Msg:  req.Msg,
	}
}

func GetMsgRequestDto2Dao(req *dto.GetMsgRequest) *dao.GetMsgRequest {
	return &dao.GetMsgRequest{
		Path: req.Path,
		Base: req.Num,
	}
}

func GetMsgResponseDao2Dto(req *dao.GetMsgResponse) *dto.GetMsgResponse {
	return &dto.GetMsgResponse{
		Meta: req.Meta,
		Msgs: req.Msgs,
	}
}

func GetMsgRequestVo2Dto(req *vo.GetMsgRequest) *dto.GetMsgRequest {
	return &dto.GetMsgRequest{
		Path: req.Path,
		Num:  req.Num,
	}
}

func GetMsgResponseDto2Vo(req *dto.GetMsgResponse) *vo.GetMsgResponse {
	return &vo.GetMsgResponse{
		Meta: req.Meta,
	}
}
