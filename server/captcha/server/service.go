package server

import (
	"net/smtp"
	"truffle/client"
	truffle_i18n "truffle/i18n/proto"
	"truffle/utils"

	"go.uber.org/zap"
)

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

func (service *CaptchaService) SendEmail(eopt *utils.EmailOpt, to []string, i18nCli truffle_i18n.I18NClient, lang string) error {
	email := utils.NewEmail()
	email.From = eopt.From
	email.To = to
	email.Subject = client.T(i18nCli, lang, "captcha.email.subject").GetMsg()
	email.Text = []byte(client.T(i18nCli, lang, "captcha.email.verification_code").GetMsg() +
		": " + utils.RandValue(utils.RandInt(100))[:6])
	err := email.Send(eopt.Server, smtp.PlainAuth(
		eopt.Auth.Id,
		eopt.Auth.Username,
		eopt.Auth.Password,
		eopt.Auth.Host,
	))
	if err != nil {
		zap.S().Warnf("Fail to send email, err: %v", err)
		return err
	}
	return nil
}
