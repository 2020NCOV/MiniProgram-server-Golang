package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// GetBindInfoService 管理用户绑定信息服务
type GetBindInfoService struct {
	UID   int    `form:"uid" json:"uid"`
	Token string `form:"token" json:"token"`
}

// GetBindInfo 检查用户绑定信息
func (service *GetBindInfoService) GetBindInfo(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	/*
		原生sql版本
	*/
	var corpCode string
	if err := model.DB2.QueryRow("select o.corp_code from wx_mp_bind_info as u left join organization as o on o.id = u.orgid where u.wx_uid = ? and u.isbind = ?", service.UID, 1).Scan(&corpCode); err != nil || corpCode == "" {
		return serializer.BuildBindInfoResponse(0, 0, "")
	}
	return serializer.BuildBindInfoResponse(0, 1, corpCode)

	/*
		gorm版本
	*/

	//var bindInfo BindIdnfo
	//
	//
	//if err := model.DB.Model(model.WxMpBindInfo{}).Select("wx_mp_bind_infos.orgid, corps.corpid, corps.corpname").Joins("left join corps on corps.id = wx_mp_bind_infos.orgid").Where(model.WxMpBindInfo{Isbind: 1, WxUid: service.UID}).First(&bindInfo).Error; err != nil {
	//	return serializer.BuildBindInfoResponse(0, 0, "")
	//}
	//bindCorpid := bindInfo.Corpid
	//return serializer.BuildBindInfoResponse(0, 1, bindCorpid)

}
