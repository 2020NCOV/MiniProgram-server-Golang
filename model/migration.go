package model

//数据库迁移
func migration() {
	//自动迁移模式
	DB.AutoMigrate(&Code{})
	DB.AutoMigrate(&DailyInfo{})
	DB.AutoMigrate(&Corp{})
	DB.AutoMigrate(&Student{})

	// 新版表
	DB.AutoMigrate(&WeChat{})
	DB.AutoMigrate(&Reporter{})
	DB.AutoMigrate(&Record{})
	//临时
	DB.AutoMigrate(&WxMpBindInfo{})
}
