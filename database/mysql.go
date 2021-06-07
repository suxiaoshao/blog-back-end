package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB = nil

func init() {
	var err error
	MysqlDb, err = getMysqlDB()
	if err != nil {
		panic(err)
	}
}

// getMysqlDB
//
// 用于获取 mysql 的数据连接接口
func getMysqlDB() (*gorm.DB, error) {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
