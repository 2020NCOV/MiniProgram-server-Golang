package api

import (
	"Miniprogram-server-Golang/service"
	"github.com/gin-gonic/gin"
)

// GetBindInfo 检查用户绑定信息接口
func GetBindInfo(c *gin.Context) {
	var service service.GetBindInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetBindInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
