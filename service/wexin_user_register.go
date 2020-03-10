package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// WeixinUserRegister CheckIsRegisteredService 管理用户注册服务
type WeixinUserRegister struct {
	UserID   string `form:"userid" json:"userid"`
	Corpid   string `form:"corpid" json:"corpid"`
	UID      string `form:"uid" json:"uid"`
	Token    string `form:"token" json:"token"`
	Name     string `form:"name" json:"name"`
	PhoneNum string `form:"phone_num" json:"phone_num"`
}

// UserRegister isRegistered 判断用户是否注册过
func (service *WeixinUserRegister) UserRegister(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	count := 0
	//在搜索数据库，判断是否存在该用户
	if model.DB.Model(&model.Student{}).Where("uid = ?", service.UID).Count(&count); count > 0 {
		return serializer.ParamErr("该用户已注册", nil)
	}

	count = 0
	//在搜索数据库，判断是否存在该用户
	if model.DB.Model(&model.Student{}).Where("user_id = ?", service.UserID).Count(&count); count > 0 {
		return serializer.ParamErr("该昵称已被占用", nil)
	}

	user := model.Student{
		Name:         service.Name,
		PhoneNum:     service.PhoneNum,
		UID:          service.UID,
		UserID:       service.UserID,
		Corpid:       service.Corpid,
		IsRegistered: 1,
		Password:     "password",
	}

	// 用户信息存库
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildIsRegisteredResponse(1)
}
