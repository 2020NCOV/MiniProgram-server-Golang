package model

//数据库迁移
func migration() {
	//自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Code{})
	DB.AutoMigrate(&DailyInfo{})
	DB.AutoMigrate(&Corp{})
	DB.AutoMigrate(&Student{})
}
