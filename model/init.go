package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(conString string) {
	db, err := gorm.Open("mysql", conString)

	if err != nil {
		panic("fail to connect database")
	}

	DB = db

	migration()
}
