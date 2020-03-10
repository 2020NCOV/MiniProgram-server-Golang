package model

import "github.com/jinzhu/gorm"

//	临时创建使用，后面按照命名规则和规定属性等再进行修改
//	对应数据库中的wx_mp_bind_info表
type WxMpBindInfo struct {
	gorm.Model
	OrgId		uint
	WxUid		string
	Username	string
	Isbind		int
}