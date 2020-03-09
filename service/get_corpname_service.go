package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// GetCorpService CheckIsRegisteredService 管理用户注册服务
type GetCorpService struct {
	UID    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
	Corpid string `form:"corpid" json:"corpid"`
}

// GetCorp isRegistered 判断用户是否注册过
func (service *GetCorpService) GetCorp(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var corp model.Corp
	if err := model.DB.Where("corpid = ?", service.Corpid).First(&corp).Error; err != nil {
		return serializer.ParamErr("无数据", nil)
	}

	return serializer.BuildCorpResponse(corp)
}
