package client

import (
	"context"
	. "truffle/ws/proto"

	"go.uber.org/zap"
)

func (c *GClientConn) NewWsClient() WsClient {
	return NewWsClient(c)
}

func WriteMsg(c WsClient, recvHandle func(*WriteMsgResponse) error) {
	reqChan := make(chan *WriteMsgRequest, 20)
	stream, err := c.WriteMsg(context.Background())
	if err != nil {
		zap.S().Warn(err)
	}
	go func() {
		for {
			select {
			case req := <-reqChan:
				stream.Send(req)
			}
		}
	}()
	for {
		data, err := stream.Recv()
		if err != nil {
			zap.S().Warn(err)
		}
		if err := recvHandle(data); err != nil {
			break
		}
	}
	return
}
