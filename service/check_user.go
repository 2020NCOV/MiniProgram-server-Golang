package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// CheckUserService CheckIsRegisteredService 管理用户注册服务
type CheckUserService struct {
	UserID string `form:"userid" json:"userid"`
	Corpid string `form:"corpid" json:"corpid"`
	UID    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
}

// CheckUser isRegistered 判断用户是否注册过
func (service *CheckUserService) CheckUser(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//再搜索数据库，判断是否存在该用户
	count := 0
	if model.DB.Model(&model.Student{}).Where(&model.Student{Uid: service.UID}).Count(&count); count == 0 {
		return serializer.BuildUserCheckResponse(0, service.Corpid, service.UserID)
	}

	return serializer.BuildUserCheckResponse(1, service.Corpid, service.UserID)
}
