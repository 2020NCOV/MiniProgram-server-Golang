package service

import (
	"github.com/gin-gonic/gin"
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
)

// CheckIsRegisteredService 管理用户注册服务
type GetInfoService struct {
	Uid    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
	Corpid string `form:"corpid" json:"corpid"`
}

// GetLastData 获取上次提交的数据
func (service *GetInfoService) GetLastData(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var lastData model.DailyInfo
	if err := model.DB.Where("uid = ?", service.Uid).First(&lastData).Error; err != nil {
		return serializer.ParamErr("无数据", nil)
	}

	return serializer.BuildLastDataResponse(lastData)
}

// GetMyInfo 获取用户数据
func (service *GetInfoService) GetMyInfo(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var user model.Student

	if err := model.DB.Where("uid = ?", service.Uid).First(&user).Error; err != nil {
		return serializer.ParamErr("无数据", nil)
	}

	return serializer.BuildUserInfoResponse(user)
}
