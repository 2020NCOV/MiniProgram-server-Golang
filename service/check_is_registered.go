package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// CheckIsRegisteredService 管理用户注册服务
type CheckIsRegisteredService struct {
	Code   string `form:"code" json:"code"`
	Corpid string `form:"corpid" json:"corpid"`
	UID    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
}

// IsRegistered 判断用户是否注册过
func (service *CheckIsRegisteredService) IsRegistered(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//到student表中找是否存在
	//在搜索数据库，判断是否存在该用户
	count := 0
	if model.DB.Model(&model.Student{}).Where(&model.Student{Uid: service.UID}).Count(&count); count == 0 {
		return serializer.BuildIsRegisteredResponse(0)
	}
	var student model.Student
	model.DB.Where("uid = ?", service.UID).First(&student)
	return serializer.BuildIsRegisteredResponse(student.IsRegistered)

}
