package server

import (
	. "truffle/i18n/proto"
	. "truffle/i18n/store"
)

type I18NService struct{}

func (s *I18NService) T(lang, msg string) *I18NResponse {
	t, _ := GetStore().Get(lang, msg)
	return &I18NResponse{Msg: t}
}

func (s *I18NService) TWithDefault(lang, msg string) *I18NResponse {
	t, ok := GetStore().Get(lang, msg)
	if !ok {
		t, ok = GetStore().Get("en_us", msg)
	}
	return &I18NResponse{Msg: t}
}
