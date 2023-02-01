package client

import (
	"context"
	. "truffle/log/proto"

	"go.uber.org/zap"
)

func (c *GClientConn) NewLogClient() LogClient {
	return NewLogClient(c)
}

func Logger(c LogClient, level, msg string) *LogResponse {
	ctx := context.Background()
	res, err := c.Logger(ctx, &LogRequest{Level: level, Msg: msg})
	if err != nil {
		zap.S().Warn(err)
	}
	return res
}
