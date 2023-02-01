package client

import (
	"context"
	. "truffle/captcha/proto"

	"go.uber.org/zap"
)

func (c *GClientConn) NewCaptchaClient() CaptchaClient {
	return NewCaptchaClient(c)
}

func SendEmail(c CaptchaClient, lang string, to []string) *EmailResponse {
	ctx := context.Background()
	res, err := c.SendEmail(ctx, &EmailRequest{To: to, Lang: lang})
	if err != nil {
		zap.S().Warn(err)
	}
	return res
}
