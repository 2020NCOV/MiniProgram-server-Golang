package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// 管理用户注册服务
type CheckIsRegisteredService struct {
	Code   string `form:"code" json:"code"`
	Corpid string `form:"corpid" json:"corpid" binding:"required"`
	UID    int    `form:"uid" json:"uid" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

// 判断用户是否注册过
func (service *CheckIsRegisteredService) IsRegistered(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//到organization表中查找是否有该企业
	var orgid string
	if err := model.DB.QueryRow("select id from organization where corp_code =?", service.Corpid).Scan(&orgid); err != nil || orgid == "" {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	////到wx_mp_bind_info表中查找是否有绑定信息
	var bindid string
	if err := model.DB.QueryRow("select id from wx_mp_bind_info where org_id =? and wx_uid =? and isbind =?", orgid, service.UID, 1).Scan(&bindid); err != nil || bindid == "" {
		return serializer.BuildIsRegisteredResponse(0, 0)
	} else {
		return serializer.BuildIsRegisteredResponse(0, 1)
	}

}
