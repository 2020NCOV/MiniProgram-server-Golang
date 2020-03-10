package serializer

import "Miniprogram-server-Golang/model"

// Status 状态
type Status struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}

// Student 用户序列化器
type Student struct {
	UID          string `json:"uid"`
	Name         string `json:"name"`
	PhoneNum     string `json:"phone_num"`
	UserID       string `json:"userid"`
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
}

// Corp 表单模板序列化器
type Corp struct {
	ErrCode      int    `json:"errcode"`
	Corpid       string `json:"corpid"`
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
	TemplateCode string `json:"template_code"`
}

// IsRegistered 用户序列化器
type IsRegistered struct {
	IsRegistered int `json:"is_registered"`
}

// CheckUser 检查
type CheckUser struct {
	IsExist int    `json:"is_exist"`
	UserID  string `json:"userid"`
	Corpid  string `json:"corpid"`
}

// BuildUserCheck 序列化
func BuildUserCheck(x int, corpid string, userid string) CheckUser {
	return CheckUser{
		IsExist: x,
		Corpid:  corpid,
		UserID:  userid,
	}
}

// BuildUserInfo 序列化
func BuildUserInfo(user model.Student) Student {
	return Student{
		UID:          user.UID,
		Name:         user.Name,
		PhoneNum:     user.PhoneNum,
		UserID:       user.UserID,
		Corpname:     user.Corpid,
		TypeCorpname: "组织编号",
		TypeUsername: "学号",
	}
}

// BuildCorp 序列化corp
func BuildCorp(errCode int, corp model.Corp) Corp {
	return Corp{
		ErrCode:      errCode,
		Corpid:       corp.Corpid,
		Corpname:     corp.Corpname,
		TypeCorpname: corp.TypeCorpname,
		TypeUsername: corp.TypeUsername,
		TemplateCode: corp.TemplateCode,
	}
}

// BuildStatus 序列化status
func BuildStatus(info model.Code) Status {
	return Status{
		UID:   info.UID,
		Token: info.Token,
	}
}

// BuildStatusResponse 序列化status响应，返回token、uid等信息
func BuildStatusResponse(info model.Code) Response {
	return Response{
		Data: BuildStatus(info),
	}
}

// BuildCorpResponse 序列化corp响应
func BuildCorpResponse(errCode int, corp model.Corp) Response {
	return Response{
		Data: BuildCorp(errCode, corp),
	}
}

// BuildIsRegisteredResponse 序列化用户注册响应
func BuildIsRegisteredResponse(x int) Response {
	return Response{
		Data: IsRegistered{IsRegistered: x},
	}
}

// BuildUserCheckResponse 序列化验证用户是否响应
func BuildUserCheckResponse(x int, corpid string, userid string) Response {
	return Response{
		Data: BuildUserCheck(x, corpid, userid),
	}
}

// BuildUserInfoResponse 序列化用户信息响应
func BuildUserInfoResponse(user model.Student) Response {
	return Response{
		Data: BuildUserInfo(user),
	}
}
