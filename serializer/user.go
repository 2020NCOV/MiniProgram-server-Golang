package serializer

import "Miniprogram-server-Golang/model"

// Status 状态
type Status struct {
	UID          string `json:"uid"`
	Token        string `json:"token"`
	IsRegistered int    `json:"is_registered"`
	ErrCode      int    `json:"errcode"`
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
	Depid        int    `json:"depid"`
}

// IsRegistered 用户序列化器
type IsRegistered struct {
	//php代码和api中还有errcode参数
	ErrCode      int `json:"errcode"`
	IsRegistered int `json:"is_registered"`
}

// CheckUser 检查
type CheckUser struct {
	ErrCode int    `json:"errcode"`
	UserID  string `json:"userid"`
	CorpID  string `json:"corpid"`
	IsExist int    `json:"is_exist"`
}

// BuildUserCheck 序列化
func BuildUserCheck(errCode int, corpID string, userID string, x int) CheckUser {
	return CheckUser{
		ErrCode: errCode,
		CorpID:  corpID,
		UserID:  userID,
		IsExist: x,
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
		Depid:        corp.Id,
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
func BuildStatusResponse(token string, uid string, isRegistered int, errcode int) Response {
	return Response{
		Token:        token,
		UID:          uid,
		IsRegistered: isRegistered,
		ErrCode:      errcode,
	}
}

// BuildCorpResponse 序列化corp响应
func BuildCorpResponse(errCode int, corp model.Corp) Response {
	return Response{
		Data: BuildCorp(errCode, corp),
	}
}

//BuildCorpResponse 序列化IsRegistered响应
func BuildIsRegistered(errcode int, is_registered int) IsRegistered {
	return IsRegistered{
		ErrCode:      errcode,
		IsRegistered: is_registered,
	}
}

// BuildIsRegisteredResponse 序列化用户注册响应
func BuildIsRegisteredResponse(errcode int, is_registered int) Response {
	return Response{
		Data: BuildIsRegistered(errcode, is_registered),
	}
}

// BuildUserCheckResponse 序列化验证用户是否响应
func BuildUserCheckResponse(errCode int, corpID string, userID string, x int) Response {
	return Response{
		Data: BuildUserCheck(errCode, corpID, userID, x),
	}
}
