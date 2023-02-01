package server

import (
	"database/sql"
	"time"
	"truffle/db"
	. "truffle/user/proto"
	"truffle/utils"

	"go.uber.org/zap"
)

type UserService struct {
	*UserMapper
}

func NewUserService() *UserService {
	return &UserService{
		UserMapper: NewUserMapper(),
	}
}

func (service *UserService) Login(key, pwd string) (ret *LoginResponse) {
	ret = new(LoginResponse)
	if len(key) <= 0 || len(pwd) <= 0 {
		ret.Msg = "login.key_pwd_null"
		return
	}
	var user User
	if utils.CheckEmail(key) {
		user = service.UserMapper.FindByEmail(key)
	} else if utils.CheckPhone(key) {
		user = service.UserMapper.FindByTelephone(key)
	} else {
		user = service.UserMapper.FindByName(key)
	}
	if user.Id == 0 {
		ret.Msg = "login.user_null"
		return
	}
	if !utils.EqualMD5(pwd, user.Password) {
		ret.Msg = "login.user_pwd_err"
		return
	}
	var token string
	var err error
	token = db.GetRedis().Get(user.Username)
	if token == "" {
		token, err = utils.ReleaseToken(service.Hash(user))
		if err != nil {
			ret.Msg = "login.token_gen_err"
			zap.S().Warn(err)
			return
		}
	} else {
		_, claims, err := utils.ParseToken(token)
		if err != nil {
			ret.Msg = "login.token_parse_err"
			zap.S().Warn(err)
			return
		}
		if utils.NowTimestamp() > claims.ExpiresAt {
			token, err = utils.ReleaseToken(service.Hash(user))
			if err != nil {
				ret.Msg = "login.token_gen_err"
				zap.S().Warn(err)
				return
			}
		}
	}
	db.GetRedis().Set(user.Username, token, time.Hour*24)
	ret.Msg = token
	ret.Ok = true
	return
}

func (service *UserService) Register(name, pwd, email string) (ret *RegisterResponse) {
	ret = new(RegisterResponse)
	if len(name) <= 0 || len(pwd) <= 0 {
		ret.Msg = "register.key_pwd_null"
		return
	}
	if len(email) <= 0 {
		ret.Msg = "register.email_null"
		return
	}
	// proof email
	// captcha

	user := User{
		Username: name,
		Password: utils.MD5(pwd),
		Email: sql.NullString{
			Valid:  true,
			String: email,
		},
	}
	if err := service.CreateUser(user); err != nil {
		ret.Msg = "register.fail"
		return
	}
	ret.Ok = true
	ret.Msg = "register.success"
	return
}
