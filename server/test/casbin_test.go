package test

import (
	"fmt"
	"testing"
	"truffle/db"
	"truffle/middleware"
	"truffle/utils"

	_ "github.com/go-sql-driver/mysql"
)

func TestCasbin(t *testing.T) {
	opt := utils.LoadOpt(utils.ENV, db.SQLOpt{})
	middleware.InitCasbin(opt.Opt(db.SQL).(db.SQLOpt))
	e := middleware.GetCasbin()
	e.AddPolicy("alice", "data1", "read")
	sub := "alice"
	obj := "data1"
	act := "read"
	// admin, err := e.AddGroupingPolicy("alice", "admin")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(admin)
	fmt.Println(middleware.CheckUserGroup("alice", "admins"))
	ok := middleware.CheckUser(sub, obj, act)
	if ok {
		fmt.Println("pass")
	} else {
		fmt.Println("unpass")
	}
	// e.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data2", "write"})
	fmt.Println(e.GetFilteredPolicy(0, "alice"))
}
