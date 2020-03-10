package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	Corpid       string
	TemplateCode string
	Corpname     string
	TypeCorpname string
	TypeUsername string
}

// User 用户模型
type User struct {
	gorm.Model
	Name     string
	Password string
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

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	//密码加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
