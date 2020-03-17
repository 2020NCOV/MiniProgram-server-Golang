package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"        //
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

// DB 适配其他版本的数据库连接实例
var DB *sql.DB

//  创建数据库连接实例
func Database(conString string) {
	db, err := sql.Open("mysql", conString)

	if err != nil {
		panic("fail to connect database")
	}

	DB = db
}
