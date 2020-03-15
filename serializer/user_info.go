package serializer

import "Miniprogram-server-Golang/model"

// UserInfo 用户数据序列化器
type UserInfo struct {
	ErrCode      int    `json:"errcode"`
	Name         string `json:"name"`
	PhoneNum     string `json:"phone_num"`
	UserID       string `json:"userid"`
	TemplateCode string `json:"bind_corp_template_code"`
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
}

// BuildUserInfo 序列化
func BuildUserInfo(user model.Student, corp model.Corp) UserInfo {
	return UserInfo{
		ErrCode:      0,
		Name:         user.Name,
		PhoneNum:     user.PhoneNum,
		UserID:       user.UserID,
		TemplateCode: corp.TemplateCode,
		Corpname:     corp.Corpname,
		TypeCorpname: corp.TypeCorpname,
		TypeUsername: corp.TypeUsername,
	}
}

// BuildUserInfoResponse 序列化用户信息响应
func BuildUserInfoResponse(user model.Student, corp model.Corp) Response {
	return Response{
		Data: BuildUserInfo(user, corp),
	}
}
