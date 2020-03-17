package conf

import (
	"Miniprogram-server-Golang/model"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置
func Init() {
	//加载.env文件
	godotenv.Load()

	//连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}
