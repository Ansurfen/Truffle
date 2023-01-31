package db

import (
	"fmt"
	"truffle/utils"
)

const (
	SQL   = "sql"
	NOSQL = "nosql"
)

type SQLOpt struct {
	Driver string
	Args   string
}

func (opt SQLOpt) Scheme() string {
	return SQL
}

func (opt SQLOpt) Init(env string, c *utils.Conf) utils.IOpt {
	database := c.GetString("db.database")
	host := c.GetString("db.host")
	port := c.GetString("db.port")
	username := c.GetString("db.username")
	password := c.GetString("db.password")
	charset := c.GetString("db.charset")
	opt.Args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	opt.Driver = c.GetString("db.driver")
	return opt
}

type NoSQLOpt struct {
	Addr string
	Pwd  string
}

func (opt NoSQLOpt) Scheme() string {
	return NOSQL
}

func (opt NoSQLOpt) Init(env string, c *utils.Conf) utils.IOpt {
	opt.Addr = c.GetString("redis.host") + ":" + c.GetString("redis.port")
	opt.Pwd = c.GetString("redis.password")
	return opt
}
