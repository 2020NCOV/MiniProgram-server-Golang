package api

import (
	"Miniprogram-server-Golang/service"

	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口，获取openid，token
func UserLogin(c *gin.Context) {
	var service service.UserOpenIDService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCode(c)
		c.JSON(200, res.Data)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserIsReg 判断用户是否存在
func UserIsReg(c *gin.Context) {
	var service service.CheckIsRegisteredService
	if err := c.ShouldBind(&service); err == nil {
		res := service.IsRegistered(c)
		if res.Data != nil {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// SaveInfo 用户上传信息接口
func SaveInfo(c *gin.Context) {
	var service service.SaveDailyInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SaveDailyInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetInfo 用户上传信息接口
func GetInfo(c *gin.Context) {
	var service service.GetLastDataService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetLastData(c)
		if res.Code == 0 {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetUserInfo 用户上传信息接口
func GetUserInfo(c *gin.Context) {
	var service service.GetInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMyInfo(c)
		if res.Code == 0 {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetCorp 获取用户企业信息接口
func GetCorp(c *gin.Context) {
	var service service.GetCorpService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCorp(c)
		if res.Data != nil {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CheckUser 检查用户是否存在
func CheckUser(c *gin.Context) {
	//	将请求的内容通过ShouldBind方法绑定到service中。每一个接口中对应的service
	var service service.CheckUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CheckUser(c)
		if res.Data != nil {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// WeixinUsrRegister 检查用户是否存在
func WeixinUsrRegister(c *gin.Context) {
	var service service.WeixinUserRegister
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserRegister(c)
		if res.Data != nil {
			c.JSON(200, res.Data)
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserUnBind 检查用户是否存在
func UserUnBind(c *gin.Context) {
	var service service.UserBindService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UnBind(c)
		c.JSON(200, res.Data)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
