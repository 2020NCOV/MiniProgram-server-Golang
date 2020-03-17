package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

//  管理用户注册服务
type CheckUserService struct {
	UserID string `form:"userid" json:"userid"`
	CorpID string `form:"corpid" json:"corpid"`
	UID    int    `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
}

// 用于检测用户标识是否已经被绑定
func (service *CheckUserService) CheckUser(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var depId int
	//	根据企业的标识码CorpID找到该企业信息对应的id
	if err := model.DB.QueryRow("select id from organization where corp_code = ?", service.CorpID).Scan(&depId); err != nil || depId == 0 {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	var wxUid int
	//	根据企业id和用户信息，查找二者的绑定情况
	err := model.DB.QueryRow("select wx_uid from wx_mp_bind_info where org_id = ? and username = ? and isbind = ?", depId, service.UserID, 1).Scan(&wxUid)

	if err == nil {
		if wxUid == service.UID {
			//	待增加bind接口后修改isExist字段值
			return serializer.BuildUserCheckResponse(0, service.CorpID, service.UserID, 0)
		} else {
			return serializer.Err(100020, "该用户已被其他微信绑定，每个用户只能被一个微信绑定", nil)
		}
	}
	return serializer.BuildUserCheckResponse(0, service.CorpID, service.UserID, 0)
}
