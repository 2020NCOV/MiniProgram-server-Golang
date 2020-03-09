package model

import (
	"github.com/jinzhu/gorm"
)

// Code 记录用户token信息
type Code struct {
	gorm.Model
	Uid   string
	Token string
	Code  string
}

// CheckToken 判断token是否正确
func CheckToken(uid string, token string) bool {
	count := 0
	if DB.Model(&Code{}).Where("uid = ? and token = ?", uid, token).Count(&count); count == 0 {
		return false
	}
	return true
}
