package server

import (
	"context"
	truffle "truffle/i18n/proto"
)

type I18NController struct {
	truffle.UnimplementedI18NServer
	*I18NService
}

func NewI18NController() *I18NController {
	return &I18NController{I18NService: &I18NService{}}
}

func (s *I18NController) T(ctx context.Context, req *truffle.I18NRequest) (*truffle.I18NResponse, error) {
	return s.TWithDefault(req.GetLang(), req.GetMsg()), nil
}

func (s *I18NController) TS(truffle.I18N_TSServer) error {
	return nil
}
