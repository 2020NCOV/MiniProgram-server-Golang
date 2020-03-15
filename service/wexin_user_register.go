package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// WeixinUserRegister CheckIsRegisteredService 管理用户注册服务
type WeixinUserRegister struct {
	UserID   string `form:"userid" json:"userid" binding:"required"`
	Corpid   string `form:"corpid" json:"corpid" binding:"required"`
	UID      string `form:"uid" json:"uid" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	PhoneNum string `form:"phone_num" json:"phone_num" binding:"required"`
}

// UserRegister isRegistered 判断用户是否注册过
func (service *WeixinUserRegister) UserRegister(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	res, _ := model.DB2.Query("select wx_uid from wx_mp_bind_info where wx_uid = ? and org_id = ? and username = ? and isbind = 1", service.UID, service.Corpid, service.Name)

	if res.Next() {
		return serializer.BuildIsRegisteredResponse(0, 1)
	} else {
		res, _ := model.DB2.Query("select org_id from wx_mp_bind_info where wx_uid = ? and username = ? and isbind = 1", service.UID, service.Name)
		if res.Next() {
			return serializer.Err(100020, "本微信已经绑定其他机构，不能重复绑定", nil)
		}
		result := model.DB2.QueryRow("insert into wx_mp_bind_info(wx_uid,org_id,username,isbind) values(?,?,?,1)", service.UID, service.Corpid, service.Name)
		if result == nil {
			return serializer.Err(50001, "注册失败", nil)
		}
		result = model.DB2.QueryRow("update wx_mp_user set userid=?, name=?, phone_num=? where wid=?", service.UserID, service.Name, service.PhoneNum, service.UID)
		if result == nil {
			return serializer.Err(50001, "用户更新失败", nil)
		}
	}
	return serializer.BuildIsRegisteredResponse(0, 1)
}
