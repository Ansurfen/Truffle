package utils

import "github.com/jordan-wright/email"

type EmailOption struct {
	From   string
	Server string
	Auth   struct {
		Id       string
		Username string
		Password string
		Host     string
	}
}

type EmailOpt struct {
	From   string
	Server string
	Auth   struct {
		Id       string
		Username string
		Password string
		Host     string
	}
}

func (opt EmailOpt) Scheme() string {
	return EMAIL
}

func (opt EmailOpt) Init(env string, c *Conf) IOpt {
	opt.From = c.GetString("email.from")
	opt.Server = c.GetString("email.server")
	opt.Auth = struct {
		Id       string
		Username string
		Password string
		Host     string
	}{
		c.GetString("email.auth.id"),
		c.GetString("email.auth.username"),
		c.GetString("email.auth.password"),
		c.GetString("email.auth.host"),
	}
	return opt
}

type Email struct {
	*email.Email
}

func NewEmail() *Email {
	return &Email{email.NewEmail()}
}

func (e *Email) Reset() {
	e.To = []string{}
}
