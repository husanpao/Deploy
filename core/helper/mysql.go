package helper

import (
	"github.com/sohaha/zlsgo/zlog"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/mysql"
	"service_manager/core/config"
)

func NewMysqlContext(c config.DataBaseConfig) *zdb.DB {
	config := &mysql.Config{
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		DBName:   c.DBName,
	}
	db, err := zdb.New(config)
	if err != nil {
		zlog.Panic(err)
	}

	return db
}
