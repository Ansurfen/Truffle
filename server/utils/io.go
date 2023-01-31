package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Conf struct {
	*viper.Viper
}

func NewConf(name, _type, dir string) *Conf {
	conf := viper.New()
	conf.SetConfigName(name)
	conf.SetConfigType(_type)
	conf.AddConfigPath(dir)
	AssertWithHandle(conf.ReadInConfig(), func(e error) {
		log.Fatal(e)
	})
	return &Conf{conf}
}

func (c *Conf) GetLogConfigs() LogConfigs {
	return LogConfigs{
		Level:       c.GetString("logger.level"),
		Format:      c.GetString("logger.format"),
		Path:        c.GetString("logger.path"),
		FileName:    c.GetString("logger.filename"),
		FileMaxSize: c.GetInt("logger.maxsize"),
		Compress:    c.GetBool("logger.compress"),
		Stdout:      c.GetBool("logger.stdout"),
	}
}
