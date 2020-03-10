package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// GetBindInfoService 请加注释
type GetBindInfoService struct {
	UID   string `form:"uid" json:"uid"`
	Token string `form:"token" json:"token"`
}

// BindInfo 请加注释
type BindInfo struct {
	Orgid    uint
	Corpid   string
	Corpname string
}

// GetBindInfo 请加注释
func (service *GetBindInfoService) GetBindInfo(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var bindInfo BindInfo
	if err := model.DB.Model(model.WxMpBindInfo{}).Select("wx_mp_bind_infos.orgid, corps.corpid, corps.corpname").Joins("left join corps on corps.id = wx_mp_bind_infos.orgid").Where(model.WxMpBindInfo{Isbind: 1, WxUid: service.UID}).First(&bindInfo).Error; err != nil {
		return serializer.BuildBindInfoResponse(0, 0, "")
	}
	bindCorpid := bindInfo.Corpid
	return serializer.BuildBindInfoResponse(0, 1, bindCorpid)

}
