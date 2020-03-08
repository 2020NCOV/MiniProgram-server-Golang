package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//用户模型
type User struct {
	gorm.Model
	Name     string
	Password string
}

type Student struct {
	gorm.Model
	Name         string
	PhoneNum     string
	Uid          string
	UserId       string
	Corpid       string
	IsRegistered int
	Password     string
}

//记录不同机构的不同模板号
type Corp struct {
	gorm.Model
	Corpid       string
	TemplateCode string
	Corpname     string
	TypeCorpname string
	TypeUsername string
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
