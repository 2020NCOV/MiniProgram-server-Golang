package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

//SaveDailyInfoService 管理每日上传信息服务

type data struct {
	IsReturnSchool            string `form:"data.is_return_school" json:"is_return_school"`
	ReturnTime                string `form:"data.return_time" json:"return_time"`
	ReturnDormNum             string `form:"data.return_dorm_num" json:"return_dorm_num"`
	ReturnTrafficInfo         string `form:"data.return_traffic_info" json:"return_traffic_info"`
	CurrentHealthValue        string `form:"data.current_health_value" json:"current_health_value"`
	CurrentContagionRiskValue string `form:"data.current_contagion_risk_value" json:"current_contagion_risk_value"`
	//ReturnDistrictValue        int       `form:"data.current_contagion_risk_value" json:"return_district_value"`
	CurrentDistrictValue int    `form:"data.current_district_value" json:"current_district_value"`
	CurrentTemperature   string `form:"data.current_temperature" json:"current_temperature"`
	PsyStatus            string `form:"data.psy_status" json:"psy_status"`
	PsyDemand            string `form:"data.psy_demand" json:"psy_demand"`
	PsyKnowledge         string `form:"data.psy_knowledge" json:"psy_knowledge"`
	Remarks              string `form:"data.remarks" json:"remarks"`
}

type SaveDailyInfoService struct {
	Data         data   `form:"data" json:"data"`
	UID          int    `form:"uid" json:"uid"`
	Token        string `form:"token" json:"token"`
	TemplateCode string `form:"template_code" json:"template_code"`
}

//判断字符串是否为空 用于解决输入数据为空 无法存储到数据库的问题
func CheckValid(s string) bool {
	if s != "" {
		return true
	}
	return false
}

// isRegistered 判断用户是否存在
func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {
	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}
	//查找用户绑定信息
	var orgid, orgname, username, userid string
	err := model.DB.QueryRow("select org_id from wx_mp_bind_info where wx_uid =? ", service.UID).Scan(&orgid)
	err = model.DB.QueryRow("select corpname from organization where id=?", orgid).Scan(&orgname)
	err = model.DB.QueryRow("select name,userid from wx_mp_user where wid =? ", service.UID).Scan(&username, &userid)
	//count 用于判断是否重复提交
	count := 0
	var time = time.Now().Format("2006-01-02")
	if service.TemplateCode == "default" {
		//学生
		//判断是否重复提交
		err = model.DB.QueryRow("select count(*) from report_record_default where userID =? and time = ?", userid, time).Scan(&count)
		if count > 0 && err == nil {
			return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
		}

		//保存信息
		queryStr := "insert into report_record_default(is_return_school,current_health_value,current_contagion_risk_value," +
			"current_district_value,current_temperature,remarks,psy_status,psy_demand,psy_knowledge,return_time," +
			"wxuid,time,org_id,org_name,name,userID,template_code,return_dorm_num,return_traffic_info)" +
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		if _, err := model.DB.Exec(queryStr, service.Data.IsReturnSchool, service.Data.CurrentHealthValue, service.Data.CurrentContagionRiskValue,
			service.Data.CurrentDistrictValue, service.Data.CurrentTemperature, service.Data.Remarks,
			service.Data.PsyStatus, service.Data.PsyDemand, service.Data.PsyKnowledge,
			sql.NullString{String: service.Data.ReturnTime, Valid: CheckValid(service.Data.ReturnTime)},
			service.UID, time, orgid, orgname, username, userid, service.TemplateCode, service.Data.ReturnDormNum,
			service.Data.ReturnTrafficInfo); err != nil {
			return serializer.ParamErr("上传失败", nil)
		}
		//企业员工
	} else {
		//判断是否重复提交
		err = model.DB.QueryRow("select count(*) from report_record_company where userID =? and time = ?", userid, time).Scan(&count)
		if count > 0 && err == nil {
			return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
		}
		//保存信息
		queryStr := "insert into report_record_company(is_return_school,current_health_value,current_contagion_risk_value," +
			"current_district_value,current_temperature,remarks,psy_status,psy_demand,psy_knowledge,return_time," +
			"wxuid,time,org_id,org_name,name,userID,template_code,return_dorm_num,return_traffic_info)" +
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		if _, err := model.DB.Exec(queryStr, service.Data.IsReturnSchool, service.Data.CurrentHealthValue, service.Data.CurrentContagionRiskValue,
			service.Data.CurrentDistrictValue, service.Data.CurrentTemperature, service.Data.Remarks,
			service.Data.PsyStatus, service.Data.PsyDemand, service.Data.PsyKnowledge,
			sql.NullString{String: service.Data.ReturnTime, Valid: CheckValid(service.Data.ReturnTime)},
			service.UID, time, orgid, orgname, username, userid, service.TemplateCode, service.Data.ReturnDormNum,
			service.Data.ReturnTrafficInfo); err != nil {
			return serializer.ParamErr("上传失败", nil)
		}
	}
	return serializer.BuildSuccessSave()
}
