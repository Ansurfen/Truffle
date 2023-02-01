package client

import (
	"context"
	. "truffle/i18n/proto"

	"go.uber.org/zap"
)

func (c *GClientConn) NewI18NClient() I18NClient {
	return NewI18NClient(c)
}

func T(c I18NClient, lang, msg string) *I18NResponse {
	ctx := context.Background()
	res, err := c.T(ctx, &I18NRequest{Lang: lang, Msg: msg})
	if err != nil {
		zap.S().Warn(err)
	}
	// if len(res.GetMsg()) <= 0 {
	// 	res.Msg = msg
	// }
	return res
}
