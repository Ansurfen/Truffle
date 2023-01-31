package service

import (
	"context"
	truffle "truffle/log/proto"
	"truffle/mq"
)

type LogController struct {
	truffle.UnimplementedLogServer
	Mq *mq.MQProducer
	*LogService
}

func NewLogController() *LogController {
	return &LogController{LogService: &LogService{}}
}

func (s *LogController) Logger(ctx context.Context, req *truffle.LogRequest) (*truffle.LogResponse, error) {
	s.LogService.PrintLogger(req.GetLevel(), req.GetMsg())
	s.LogService.PushMQ("log", req.GetLevel(), req.GetMsg(), s.Mq)
	return &truffle.LogResponse{Ok: true}, nil
}
