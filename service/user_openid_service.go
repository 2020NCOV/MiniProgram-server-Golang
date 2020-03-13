package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

// UserOpenIDService 获取用户token服务
type UserOpenIDService struct {
	Code string `form:"code" json:"code"`
}

// GetCode 用户登录函数，获取openidhesessionkey，作为之后操作的验证
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

	info := model.Code{
		UID:   res.OpenID,
		Token: res.SessionKey,
		Code:  service.Code,
	}

	//查看数据库中是否已有token信息
	err = model.DB2.QueryRow("select openid, token from wx_mp_user where openid = ?,  token  = ?", res.OpenID, res.SessionKey).
		Scan(&res.OpenID, &res.SessionKey)
	if err != nil {
		return serializer.Err(1008, "获取请求失败，请退出重试", nil)
	}
	//如果没有，重新存入并返回
	err = model.DB2.QueryRow("insert into wx_mp_user(openid, token)values(?,?)", res.OpenID, res.SessionKey)
	if err != nil {
		return serializer.BuildStatusResponse(res.SessionKey, res.OpenID, 1, 0)
	}

	return serializer.BuildStatusResponse(info)
}








