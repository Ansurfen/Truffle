package client

import (
	"context"
	. "truffle/user/proto"

	"go.uber.org/zap"
)

func (c *GClientConn) NewUserClient() UserClient {
	return NewUserClient(c)
}

func Login(c UserClient, key, pwd string) *LoginResponse {
	ctx := context.Background()
	res, err := c.Login(ctx, &LoginRequest{Key: key, Password: pwd})
	if err != nil {
		zap.S().Warn(err)
	}
	return res
}

func Register(c UserClient, lang, username, pwd, email string) *RegisterResponse {
	ctx := context.Background()
	res, err := c.Register(ctx, &RegisterRequest{
		Lang:     lang,
		Username: username,
		Password: pwd,
		Email:    email,
	})
	if err != nil {
		zap.S().Warn(err)
	}
	return res
}

func Logout(c UserClient, username, jwt string) *LogoutResponse {
	ctx := context.Background()
	res, err := c.Logout(ctx, &LogoutRequest{
		Username: username,
		Jwt:      jwt,
	})
	if err != nil {
		zap.S().Warn(err)
	}
	return res
}
