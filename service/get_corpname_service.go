package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// 管理用户企业身份服务
type GetCorpService struct {
	Uid    int    `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
	Corpid string `form:"corpid" json:"corpid"`
}

// 获取用户企业信息
func (service *GetCorpService) GetCorp(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var corp model.Corp
	err := model.DB.QueryRow("select id,corp_code,corpname,template_code,type_corpname,type_username from organization where corp_code =?", service.Corpid).
		Scan(&corp.Id, &corp.Corpid, &corp.Corpname, &corp.TemplateCode, &corp.TypeCorpname, &corp.TypeUsername)
	if err != nil {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	return serializer.BuildCorpResponse(0, corp)
}
