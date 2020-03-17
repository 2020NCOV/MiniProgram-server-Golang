package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// 管理用户绑定信息服务
type GetBindInfoService struct {
	UID   int    `form:"uid" json:"uid"`
	Token string `form:"token" json:"token"`
}

// 检查用户绑定信息
func (service *GetBindInfoService) GetBindInfo(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var corpCode string
	if err := model.DB.QueryRow("select o.corp_code from wx_mp_bind_info as u left join organization as o on o.id = u.orgid where u.wx_uid = ? and u.isbind = ?", service.UID, 1).Scan(&corpCode); err != nil || corpCode == "" {
		return serializer.BuildBindInfoResponse(0, 0, "")
	}
	return serializer.BuildBindInfoResponse(0, 1, corpCode)
}
