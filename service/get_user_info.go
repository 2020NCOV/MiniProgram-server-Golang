package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetInfoService 管理获取用户数据服务
type GetInfoService struct {
	UID    int    `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
	Corpid string `form:"corpid" json:"corpid"`
}

// GetMyInfo 获取用户数据
func (service *GetInfoService) GetMyInfo(c *gin.Context) serializer.Response {
	// 处理参数错误
	if service.Corpid == "" {
		return serializer.Err(1003, "参数错误:corpid", nil)
	}
	if !model.CheckToken(strconv.Itoa(service.UID), service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	// 获取传递参数的企业信息
	var corp model.Corp
	err := model.DB2.QueryRow("select corpname,template_code,type_corpname,type_username from organization where corp_code = ?", service.Corpid).
		Scan(&corp.Corpname, &corp.TemplateCode, &corp.TypeCorpname, &corp.TypeUsername)
	if err != nil {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	// 获取用户信息
	var user model.Student
	err = model.DB2.QueryRow("select userid,name,phone_num from wx_mp_user where wid = ?", service.UID).
		Scan(&user.UserID, &user.Name, &user.PhoneNum)
	if err != nil {
		return serializer.Err(1005, "获取用户信息失败", nil)
	}

	return serializer.BuildUserInfoResponse(user, corp)
}
