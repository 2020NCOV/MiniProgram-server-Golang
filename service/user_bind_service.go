package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"github.com/gin-gonic/gin"
)

// UserBindService 管理用户注册服务
type UserBindService struct {
	UserID   string `form:"userid" json:"userid"`
	Corpid   string `form:"corpid" json:"corpid"`
	UID      string `form:"uid" json:"uid"`
	Token    string `form:"token" json:"token"`
	Password string `form:"password" json:"password"`
}

// Bind 用户绑定
func (service *UserBindService) Bind(c *gin.Context) serializer.IsRegisteredBindResponse {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.BuildIsRegisteredBindResponse(serializer.CodeParamErr,0,"token验证错误")
	}

	//在搜索数据库，判断是否存在该用户
	count := 0
	if model.DB.Model(&model.Student{}).Where(&model.Student{UID: service.UID, Password: service.Password}).Count(&count); count == 0 {
		return serializer.BuildIsRegisteredBindResponse(0,1, "不存在该用户")
	}
	var student model.Student
	model.DB.Where("uid = ?", service.UID).First(&student)
	student.IsRegistered = 2
	model.DB.Save(&student)

	return serializer.BuildIsRegisteredBindResponse(0,2,"绑定成功")
}

// UnBind 用户绑定
func (service *UserBindService) UnBind(c *gin.Context) serializer.IsRegisteredBindResponse {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.BuildIsRegisteredBindResponse(serializer.CodeParamErr,0,"token验证错误")
	}
	//在搜索数据库，修改注册状态
	var student model.Student
	//判断是否存在注册信息，按照 bind 逻辑应该 is_register = 1 提醒用户注册，但小程序未做此处理，且除了调试应该不会出现未绑定情况下请求 unbind , 因此此处返回错误码
	if model.DB.Where("uid = ?", service.UID).First(&student).RecordNotFound() {
		return serializer.BuildIsRegisteredBindResponse(serializer.CodeParamErr,1,"不存在该用户")
	}
	return serializer.BuildIsRegisteredBindResponse(0,0,"解绑成功")
}
