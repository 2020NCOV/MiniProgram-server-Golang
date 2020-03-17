package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

// 管理用户注册服务
type UserBindService struct {
	UserID   string `form:"userid" json:"userid"`
	Corpid   string `form:"corpid" json:"corpid"`
	UID      int    `form:"uid" json:"uid"`
	Token    string `form:"token" json:"token"`
	Password string `form:"password" json:"password"`
}

// 用户绑定
func (service *UserBindService) UnBind(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		// Token 验证失败 此处 is_registered 无意义
		return serializer.BuildIsRegisteredResponse(serializer.CodeParamErr, 0)
	}
	// 再搜索数据库，修改注册状态
	var isRegistered int
	err := model.DB.QueryRow("SELECT isbind FROM wx_mp_bind_info WHERE wx_uid = ?", service.UID).Scan(&isRegistered)
	if err != nil {
		// 无该用户 根据 PHP 代码此处返回 errcode: 0 is_registered: 0
		if err == sql.ErrNoRows {
			return serializer.BuildIsRegisteredResponse(0, 0)
		} else {
			// 此处为其它数据库错误 小程序并未做针对处理 此处 is_registered 无意义
			return serializer.BuildIsRegisteredResponse(serializer.CodeDBError, 0)
		}
	}
	rows, err := model.DB.Exec("UPDATE wx_mp_bind_info SET isbind = ? , unbind_date = ? WHERE wx_uid = ?;", "0", time.Now(), service.UID)
	_ = rows
	if err != nil {
		// 此处为其它数据库错误 小程序并未做针对处理 此处 is_registered 无意义
		return serializer.BuildIsRegisteredResponse(serializer.CodeDBError, 0)
	}
	return serializer.BuildIsRegisteredResponse(0, 0)
}
