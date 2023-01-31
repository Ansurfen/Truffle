package server

import (
	"context"
	truffle "truffle/captcha/proto"
	truffle_i18n "truffle/i18n/proto"
	truffle_log "truffle/log/proto"
	"truffle/utils"

	"go.uber.org/zap"
)

type CaptchaController struct {
	truffle.UnimplementedCaptchaServer
	eopt    *utils.EmailOpt
	i18nCli truffle_i18n.I18NClient
	logCli  truffle_log.LogClient
	service *CaptchaService
}

func NewCaptchaController(i18nCli truffle_i18n.I18NClient, logCli truffle_log.LogClient, opt utils.EmailOpt) *CaptchaController {
	return &CaptchaController{
		i18nCli: i18nCli,
		logCli:  logCli,
		eopt:    &opt,
		service: NewCaptchaService(),
	}
}

func (con *CaptchaController) SendEmail(ctx context.Context, req *truffle.EmailRequest) (*truffle.EmailResponse, error) {
	err := con.service.SendEmail(con.eopt, req.GetTo(), con.i18nCli, req.GetLang())
	if err != nil {
		zap.S().Warnf("Fail to send email, err: %v", err)
		return &truffle.EmailResponse{Ok: false}, nil
	}
	return &truffle.EmailResponse{Ok: true}, nil
}
