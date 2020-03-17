package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

// 获取用户token服务
type UserOpenIDService struct {
	Code string `form:"code" json:"code"`
}

// 用户登录函数，获取openid和sessionkey，作为之后操作的验证
func (service *UserOpenIDService) GetCode(c *gin.Context) serializer.Response {
	res, err := weapp.Login(os.Getenv("APP_ID"), os.Getenv("APP_SECRET"), service.Code)

	if err != nil {
		//处理错误
		return serializer.ParamErr("获取openid失败", err)
	}

	if err := res.GetResponseError(); err != nil {
		//处理小程序传送的错误信息
		return serializer.ParamErr("小程序报错", err)
	}

	//查看数据库中是否已有token信息
	var wid int64
	var token string
	//err = model.DB.QueryRow("select wid from wx_mp_user where wid = ?", UID).Scan(&wid)
	err = model.DB.QueryRow("select wid, token from wx_mp_user where openid = ?", res.OpenID).
		Scan(&wid, &token)

	if err != nil {
		//如果没有，重新存入并返回
		result, err2 := model.DB.Exec("insert into wx_mp_user(openid, token) values(?,?)", res.OpenID, res.SessionKey)
		var err3 error
		wid, err3 = result.LastInsertId()
		if err2 != nil || err3 != nil {
			return serializer.Err(1008, "获取请求失败，请退出重试", nil)
		}
	}
	return serializer.BuildStatusResponse(token, wid, 1, 0)
}
