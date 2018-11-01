package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
	"git.jiaxianghudong.com/go/logs"
)

var SqlDB *gorm.DB
func InitDB(addr string,maxOpen int,maxIdle int) error{
	var err error
	SqlDB, err = gorm.Open("mysql", addr)
	if err != nil {
		logs.Error(err)
		return err
	}
	SqlDB.LogMode(true)
	SqlDB.DB().SetMaxOpenConns(maxOpen)
	SqlDB.DB().SetMaxIdleConns(maxIdle)
	SqlDB.DB().SetConnMaxLifetime(60*time.Second)
	SqlDB.DB().Ping()
	//SqlDB.
	return nil
}
