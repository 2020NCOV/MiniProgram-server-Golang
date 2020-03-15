package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//SaveDailyInfoService 管理每日上传信息服务
type data struct {
	IsReturnSchool            string    `form:"data.is_return_school" json:"is_return_school"`
	ReturnTime                string 	`form:"data.return_time" json:"return_time"`
	ReturnDormNum             string 	`form:"data.return_dorm_num" json:"return_dorm_num"`
	ReturnTrafficInfo         string 	`form:"data.return_traffic_info" json:"return_traffic_info"`
	CurrentHealthValue        string    `form:"data.current_health_value" json:"current_health_value"`
	CurrentContagionRiskValue string	`form:"data.current_contagion_risk_value" json:"current_contagion_risk_value"`
	ReturnDistrictValue       int       `form:"data.return_district_value" json:"return_district_value"`
	CurrentDistrictValue      int       `form:"data.current_district_value" json:"current_district_value"`
	CurrentTemperature        string    `form:"data.current_temperature" json:"current_temperature"`
	PsyStatus                 string    `form:"data.psy_status" json:"psy_status"`
	PsyDemand                 string    `form:"data.psy_demand" json:"psy_demand"`
	PsyKnowledge              string    `form:"data.psy_knowledge" json:"psy_knowledge"`
	Remarks                   string    `form:"data.remarks" json:"remarks"`
	PlanCompanyDate           string    `form:"data.plan_company_date" json:"plan_company_date"`
}

type SaveDailyInfoService struct {
	Data                      data   `form:"data" json:"data"`
	UID                       int    `form:"uid" json:"uid"`
	Token                     string `form:"token" json:"token"`
	TemplateCode              string `form:"template_code" json:"template_code"`
}

//sql 版

// isRegistered 判断用户是否存在
func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {
		if !model.CheckToken(strconv.Itoa(service.UID), service.Token) {
			return serializer.ParamErr("token验证错误", nil)
		}
		//调用了其他接口的数据，做单接口测试时注释掉此段
		//var orgname string
		//err:=model.DB2.QueryRow("select corpname from organization where id=?",orgid).Scan(&orgname)
		//username :=user.Name
		//userid := user.UserID

		//单接口测试数据
		orgname:="北京大学"
		orgid :="1"
		username :="bakaka"
		userid := "123"

	   //判断是否重复提交
		var count int
		err :=model.DB2.QueryRow("select count(*) from report_record_school_df where userID =? and time = ?",userid, time.Now().Format("2006-01-02")).Scan(&count)
		if count > 0 && err==nil {
			return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
		}

		//保存信息
		var time = time.Now().Format("2006-01-02")
		if(service.TemplateCode == "school_df"){
			//保存学生信息
			queryStr := "insert into report_record_school_df(is_return_school,current_health_value,current_contagion_risk_value,return_district_value," +
				"current_district_value,current_temperature,remarks,psy_status,psy_demand,psy_knowledge,plan_company_date,return_dorm_num,return_time," +
				"return_traffic_info,wxuid,time,org_id,org_name,name,userID)" + "values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
			if _,err:= model.DB2.Exec(queryStr, service.Data.IsReturnSchool, service.Data.CurrentHealthValue, service.Data.CurrentContagionRiskValue,
				service.Data.ReturnDistrictValue, service.Data.CurrentDistrictValue, service.Data.CurrentTemperature, service.Data.Remarks,
				service.Data.PsyStatus, service.Data.PsyDemand, service.Data.PsyKnowledge, service.Data.PlanCompanyDate, service.Data.ReturnDormNum, service.Data.ReturnTime,
				service.Data.ReturnTrafficInfo, service.UID,time,orgid,orgname,username,userid); err != nil {
				return serializer.ParamErr("上传失败", nil)
				}
				//保存企业员工信息
		} else{
			queryStr := "insert into report_record_company_df(is_return_school,current_health_value,current_contagion_risk_value,return_district_value," +
				"current_district_value,current_temperature,remarks,psy_status,psy_demand,psy_knowledge,plan_company_date,return_dorm_num,return_time," +
				"return_traffic_info,wxuid,time,org_id,org_name,name,userID)" + "values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
			if _,err:= model.DB2.Exec(queryStr, service.Data.IsReturnSchool, service.Data.CurrentHealthValue, service.Data.CurrentContagionRiskValue,
				service.Data.ReturnDistrictValue, service.Data.CurrentDistrictValue, service.Data.CurrentTemperature, service.Data.Remarks,
				service.Data.PsyStatus, service.Data.PsyDemand, service.Data.PsyKnowledge, service.Data.PlanCompanyDate, service.Data.ReturnDormNum, service.Data.ReturnTime,
				service.Data.ReturnTrafficInfo, service.UID,time,orgid,orgname,username,userid); err != nil {
				return serializer.ParamErr("上传失败", nil)
			}
		}
		return serializer.BuildSuccessSave()
	}

	// Gorm版

	////isRegistered 判断用户是否存在
	//func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {
	//	if !model.CheckToken(strconv.Itoa(service.Uid), service.Token) {
	//		return serializer.ParamErr("token验证错误", nil)
	//	}
	//	dailyInfo := model.Record{
	//		IsReturnSchool:            service.IsReturnSchool,
	//		CurrentHealthValue:        service.CurrentHealthValue,
	//		CurrentContagionRiskValue: service.CurrentContagionRiskValue,
	//		ReturnDistrictValue:       service.ReturnDistrictValue,
	//		CurrentDistrictValue:      service.CurrentDistrictValue,
	//		CurrentTemperature:        service.CurrentTemperature,
	//		Remarks:                   service.Remarks,
	//		PsyStatus:                 service.PsyStatus,
	//		PsyDemand:                 service.PsyDemand,
	//		PsyKnowledge:              service.PsyKnowledge,
	//		PlanCompanyDate:           service.PlanCompanyDate,
	//		ReturnDormNum:             service.ReturnDormNum,
	//		ReturnTime:                service.ReturnTime,
	//		ReturnTrafficInfo:         service.ReturnTrafficInfo,
	//		Uid:                       service.Uid,
	//		SaveDate:                  time.Now().Format("2006-01-02"),
	//	}
	//	判断该用户这天是否已经提交过
	//	count := 0
	//	if model.DB.Model(&model.DailyInfo{}).Where("uid = ? and save_date = ?", service.Uid, time.Now().Format("2006-01-02")).Count(&count); count > 0 {
	//		return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
	//	}
	//	记录用户当日信息
	//	if err := model.DB.Create(&dailyInfo).Error; err != nil {
	//		return serializer.ParamErr("上传失败", err)
	//	}
	//	return serializer.BuildSuccessSave()
	//}
