package main

import (
	"ncov_go/conf"
	"ncov_go/server"
)

func main() {
	//从配置文件读取配置
	conf.Init()

	//装载路由
	router := server.NewRouter()

	//监听8080端口
	router.Run(":8080")
}
