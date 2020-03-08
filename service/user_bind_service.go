package service

import (
	"github.com/gin-gonic/gin"
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
)

// UserBindService 管理用户注册服务
type UserBindService struct {
	UserId   string `form:"userid" json:"userid"`
	Corpid   string `form:"corpid" json:"corpid"`
	Uid      string `form:"uid" json:"uid"`
	Token    string `form:"token" json:"token"`
	Password string `form:"password" json:"password"`
}

// Bind 用户绑定
func (service *UserBindService) Bind(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//在搜索数据库，判断是否存在该用户
	count := 0
	if model.DB.Model(&model.Student{}).Where(&model.Student{Uid: service.Uid, Password: service.Password}).Count(&count); count == 0 {
		return serializer.BuildIsRegisteredResponse(0)
	} else {
		var student model.Student
		model.DB.Where("uid = ?", service.Uid).First(&student)
		student.IsRegistered = 2
		model.DB.Save(&student)
	}

	return serializer.BuildIsRegisteredResponse(2)
}

// UnBind 用户绑定
func (service *UserBindService) UnBind(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//在搜索数据库，修改注册状态
	var student model.Student
	model.DB.Where("uid = ?", service.Uid).First(&student)
	student.IsRegistered = 1
	model.DB.Save(&student)

	return serializer.BuildIsRegisteredResponse(1)
}
