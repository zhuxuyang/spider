package resource

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB 初始化 MySQL 链接
func InitDB(user, password, host, port, dbName string) {
	mdb, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, dbName))
	if err != nil {
		panic(err)
		return
	}
	if mdb == nil {
		panic("failed to connect database")
	}

	mdb.LogMode(true)
	log.Println("connected")
	db = mdb
	return
}

// GetDB 获取数据库链接实例
func GetDB() *gorm.DB {
	return db
}
