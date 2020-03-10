package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

// DB 数据库连接实例
var DB *gorm.DB

// Database 创建数据库连接实例
func Database(conString string) {
	db, err := gorm.Open("mysql", conString)

	if err != nil {
		panic("fail to connect database")
	}

	DB = db

	migration()
}
