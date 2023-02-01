package server

import (
	"context"
	"truffle/client"
	truffle_i18n "truffle/i18n/proto"
	truffle_log "truffle/log/proto"
	truffle "truffle/user/proto"
)

type UserController struct {
	truffle.UnimplementedUserServer
	*UserService
	i18nCli truffle_i18n.I18NClient
	logCli  truffle_log.LogClient
}

func NewUserController(i18nCli truffle_i18n.I18NClient, logCli truffle_log.LogClient) *UserController {
	return &UserController{
		i18nCli:     i18nCli,
		logCli:      logCli,
		UserService: NewUserService(),
	}
}

func (con *UserController) Login(ctx context.Context, req *truffle.LoginRequest) (*truffle.LoginResponse, error) {
	res := con.UserService.Login(req.GetKey(), req.GetPassword())
	tres := client.T(con.i18nCli, "en_us", res.GetMsg())
	res.Msg = tres.Msg
	return res, nil
}

func (con *UserController) Register(ctx context.Context, req *truffle.RegisterRequest) (*truffle.RegisterResponse, error) {
	res := con.UserService.Register(req.GetUsername(), req.GetPassword(), req.GetEmail())
	tres := client.T(con.i18nCli, req.GetLang(), res.GetMsg())
	res.Msg = tres.Msg
	return res, nil
}

func (con *UserController) RegAndLogin(ctx context.Context, req *truffle.RegisterRequest) (*truffle.LoginResponse, error) {
	res := con.UserService.Register(req.GetUsername(), req.GetPassword(), req.GetEmail())
	if !res.Ok {
		return &truffle.LoginResponse{Ok: false, Msg: res.GetMsg()}, nil
	}
	ret := con.UserService.Login(req.GetUsername(), req.GetPassword())
	return ret, nil
}
