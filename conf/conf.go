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
	model.Database(os.Getenv("MYSQL_DSN"), os.Getenv("MYSQL_DSN2"))

	//判断该用户这天是否已经提交过
	count := 0
	if model.DB.Model(&model.Corp{}).Where("corpid = 100000001").Count(&count); count == 0 {
		model.DB.Create(&model.Corp{
			Corpid:       "100000001",
			Corpname:     "北京大学",
			TypeCorpname: "学校名称",
			TypeUsername: "学号",
		})
	}
}
