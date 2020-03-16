package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// CheckUserService CheckIsRegisteredService 管理用户注册服务
type CheckUserService struct {
	UserID string `form:"userid" json:"userid"`
	CorpID string `form:"corpid" json:"corpid"`
	UID    int    `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
}

// CheckUser 用于检测用户标识是否已经被绑定
func (service *CheckUserService) CheckUser(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	/*
		原生sql版本
	*/

	var depId int
	//	根据企业的标识码CorpID找到该企业信息对应的id
	if err := model.DB2.QueryRow("select id from organization where corp_code = ?", service.CorpID).Scan(&depId); err != nil || depId == 0 {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	var wxUid int
	//	根据企业id和用户信息，查找二者的绑定情况
	err := model.DB2.QueryRow("select wx_uid from wx_mp_bind_info where org_id = ? and username = ? and isbind = ?", depId, service.UserID, 1).Scan(&wxUid)

	if err == nil {
		if wxUid == service.UID {
			//	待增加bind接口后修改isExist字段值
			return serializer.BuildUserCheckResponse(0, service.CorpID, service.UserID, 0)
		} else {
			return serializer.Err(100020, "该用户已被其他微信绑定，每个用户只能被一个微信绑定", nil)
		}
	}
	return serializer.BuildUserCheckResponse(0, service.CorpID, service.UserID, 0)

	/*
		gorm版本
	*/
	////	根据corpid找到公司名称
	//var corp model.Corp
	//if err := model.DB.Where(&model.Corp{CorpID: service.CorpID}).First(&corp).Error; err != nil {
	//	return serializer.Err(10006, "获取企业信息失败", nil)
	//}
	//
	//corpID := corp.ID
	////	根据corpid查找用户-企业绑定信息
	//var corpBind model.WxMpBindInfo
	//if err := model.DB.Where(&model.WxMpBindInfo{OrgId: corpID, Username: service.UserID, Isbind: 1}).First(&corpBind).Error; err != nil {
	//	//	错误码未知，张老师没有写到，有待修改
	//	return serializer.Err(100019, "用户和企业未绑定", nil)
	//}
	//
	//wxuid := corpBind.WxUid
	//
	//if wxuid == service.UID {
	//	//	正确的返回结果
	//	return serializer.BuildUserCheckResponse(0, service.Corpid, service.UserID)
	//}
	////	这里不确定是返回错误信息还是显示用户已存在。接口文档和php代码不一致，目前以php代码为准
	//return serializer.Err(100020, "该用户已被其他微信绑定，每个用户只能被一个微信绑定", nil)
	////return serializer.BuildUserCheckResponse(1, service.Corpid, service.UserId)

}
