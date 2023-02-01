package test

import (
	"testing"
	. "truffle/db"
	. "truffle/utils"
)

type TestModel struct {
	Title      string
	Descriptor string
}

func (TestModel) TableName() string {
	return "test_tb"
}

func TestSQL(t *testing.T) {
	opt := LoadOpt(ENV_DEVELOP, DefaultOpt{}, SQLOpt{})
	InitDB(opt.Opt(SQL).(SQLOpt))
	db := GetDB()
	db.AutoMigrate(&TestModel{})
}
