package store

import "truffle/utils"

const LANG = "lang"

type LangOpt struct {
	Langs []string
}

func (opt LangOpt) Scheme() string {
	return LANG
}

func (opt LangOpt) Init(env string, c *utils.Conf) utils.IOpt {
	opt.Langs = c.GetStringSlice("langs")
	return opt
}
