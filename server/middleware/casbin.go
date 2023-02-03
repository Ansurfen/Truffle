package middleware

import (
	"strings"
	"truffle/db"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var e *casbin.Enforcer

func InitCasbin(opt db.SQLOpt) {
	a, _ := gormadapter.NewAdapter(opt.Driver, opt.Args, true)
	e, _ = casbin.NewEnforcer("./model.conf", a)
	e.AddFunction("truffle_func", TruffleMatchFunc)
}

func CheckUserGroup(user, group string) bool {
	return e.HasGroupingPolicy(user, group)
}

func CheckUser(sub, obj, act string) bool {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		zap.S().Warn(err)
		return false
	}
	return ok
}

func GetCasbin() *casbin.Enforcer {
	return e
}

func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func TruffleMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
