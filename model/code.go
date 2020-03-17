package model

//  记录用户token信息
type Code struct {
	UID   string
	Token string
	Code  string
}

//  判断token是否正确
func CheckToken(uid int, token string) bool {
	res, _ := DB.Query("select wid from wx_mp_user where wid = ? and token = ?", uid, token)

	if !res.Next() {
		return false
	}

	return true
}
