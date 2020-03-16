package model

// Code 记录用户token信息
type Code struct {
	UID   string
	Token string
	Code  string
}

// CheckToken 判断token是否正确
func CheckToken(uid string, token string) bool {
	res, _ := DB2.Query("select wid from wx_mp_user where wid = ? and token = ?", uid, token)

	if !res.Next() {
		return false
	}

	return true
}
