package serializer

import "Miniprogram-server-Golang/model"

// 状态
type Status struct {
	UID          int64  `json:"uid"`
	Token        string `json:"token"`
	IsRegistered int    `json:"is_registered"`
	ErrCode      int    `json:"errcode"`
}

// 表单模板序列化器
type Corp struct {
	ErrCode      int    `json:"errcode"`
	Corpid       string `json:"corpid"`
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
	TemplateCode string `json:"template_code"`
	Depid        int    `json:"depid"`
}

// 用户序列化器
type IsRegistered struct {
	//php代码和api中还有errcode参数
	ErrCode      int `json:"errcode"`
	IsRegistered int `json:"is_registered"`
}

// 检查
type CheckUser struct {
	ErrCode int    `json:"errcode"`
	UserID  string `json:"userid"`
	CorpID  string `json:"corpid"`
	IsExist int    `json:"is_exist"`
}

// 序列化
func BuildUserCheck(errCode int, corpID string, userID string, x int) CheckUser {
	return CheckUser{
		ErrCode: errCode,
		CorpID:  corpID,
		UserID:  userID,
		IsExist: x,
	}
}

// 序列化corp
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

// 序列化status
func BuildStatus(token string, uid int64, isRegistered int, errcode int) Status {
	return Status{
		Token:        token,
		UID:          uid,
		IsRegistered: isRegistered,
		ErrCode:      errcode,
	}
}

// 序列化status响应，返回token、uid等信息
func BuildStatusResponse(token string, uid int64, isRegistered int, errcode int) Response {
	return Response{
		Data: BuildStatus(token, uid, isRegistered, errcode),
	}
}

// 序列化corp响应
func BuildCorpResponse(errCode int, corp model.Corp) Response {
	return Response{
		Data: BuildCorp(errCode, corp),
	}
}

// 序列化IsRegistered响应
func BuildIsRegistered(errcode int, is_registered int) IsRegistered {
	return IsRegistered{
		ErrCode:      errcode,
		IsRegistered: is_registered,
	}
}

//  序列化用户注册响应
func BuildIsRegisteredResponse(errcode int, is_registered int) Response {
	return Response{
		Data: BuildIsRegistered(errcode, is_registered),
	}
}

//  序列化验证用户是否响应
func BuildUserCheckResponse(errCode int, corpID string, userID string, x int) Response {
	return Response{
		Data: BuildUserCheck(errCode, corpID, userID, x),
	}
}
