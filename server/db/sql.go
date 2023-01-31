package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type SQLClient struct {
	*gorm.DB
}

var sqldb *SQLClient

func GetDB() *SQLClient {
	return sqldb
}

func InitDB(opt SQLOpt) {
	db, err := gorm.Open(opt.Driver, opt.Args)
	if err != nil {
		zap.S().Fatal(err)
	}
	err = db.DB().Ping()
	if err != nil {
		zap.S().Fatal(err)
	}
	zap.S().Info("Successfully connected...")
	sqldb = &SQLClient{
		DB: db,
	}
}
