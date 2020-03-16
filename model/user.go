package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// WeChat means Who has loged in
type WeChat struct {
	gorm.Model
	OpenID string `gorm:"unique;"`
}

// Reporter 上报人
type Reporter struct {
	gorm.Model
	WeChat         WeChat     `gorm:"association_foreignkey:WeChatRefer;unique;type:varchar(200)"`
	WeChatRefer    string     //对应的微信号
	Corp           Corp       `gorm:"foreign_key:OrgID;unique;"` //对应旧版表中的Corp
	OrgID          int        // 机构id
	Name           string     `gorm:"varchar(30);"`
	Gender         int        //0代表女性，1代表男性
	Tel            int64      //手机号
	StuNum         int64      //学号
	LastUpdateTime *time.Time `gorm:"default:null;"` //最后更新时间
}

/*
-------------------------以下为旧版表------------------------------
*/

// Corp 记录不同机构的不同模板号
type Corp struct {
	gorm.Model
	Id           int
	Corpid       string
	TemplateCode string
	Corpname     string
	TypeCorpname string
	TypeUsername string
}

// Student 学生
type Student struct {
	gorm.Model
	Name         string
	PhoneNum     string
	UID          string
	UserID       string
	Corpid       string
	IsRegistered int
	Password     string
}
