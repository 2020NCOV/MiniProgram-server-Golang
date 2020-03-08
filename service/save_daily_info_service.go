package service

import (
	"github.com/gin-gonic/gin"
	"ncov_go/model"
	"ncov_go/serializer"
	"time"
)

// SaveDailyInfoService 管理每日上传信息服务
type SaveDailyInfoService struct {
	IsReturnSchool            string `form:"is_return_school" json:"is_return_school"`
	ReturnTime                string `form:"return_time" json:"return_time"`
	ReturnDormNum             string `form:"return_dorm_num" json:"return_dorm_num"`
	ReturnTrafficInfo         string `form:"return_traffic_info" json:"return_traffic_info"`
	CurrentHealthValue        string `form:"current_health_value" json:"current_health_value"`
	CurrentContagionRiskValue string `form:"current_contagion_risk_value" json:"current_contagion_risk_value"`
	ReturnDistrictValue       string `form:"return_district_value" json:"return_district_value"`
	CurrentDistrictValue      string `form:"current_district_value" json:"current_district_value"`
	CurrentTemperature        string `form:"current_temperature" json:"current_temperature"`
	PsyStatus                 string `form:"psy_status" json:"psy_status"`
	PsyDemand                 string `form:"psy_demand" json:"psy_demand"`
	Remarks                   string `form:"remarks" json:"remarks"`
	PsyKnowledge              string `form:"psy_knowledge" json:"psy_knowledge"`
	PlanCompanyDate           string `form:"plan_company_date" json:"plan_company_date"`
	Uid                       string `form:"uid" json:"uid"`
	Token                     string `form:"token" json:"token"`
	TemplateCode              string `form:"template_code" json:"template_code"`
}

// isRegistered 判断用户是否存在
func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	dailyInfo := model.DailyInfo{
		IsReturnSchool:            service.IsReturnSchool,
		CurrentHealthValue:        service.CurrentHealthValue,
		CurrentContagionRiskValue: service.CurrentContagionRiskValue,
		ReturnDistrictValue:       service.ReturnDistrictValue,
		CurrentDistrictValue:      service.CurrentDistrictValue,
		CurrentTemperature:        service.CurrentTemperature,
		Remarks:                   service.Remarks,
		PsyStatus:                 service.PsyStatus,
		PsyDemand:                 service.PsyDemand,
		PsyKnowledge:              service.PsyKnowledge,
		PlanCompanyDate:           service.PlanCompanyDate,
		ReturnDormNum:             service.ReturnDormNum,
		ReturnTime:                service.ReturnTime,
		ReturnTrafficInfo:         service.ReturnTrafficInfo,
		Uid:                       service.Uid,
		SaveDate:                  time.Now().Format("2006-01-02"),
	}

	//判断该用户这天是否已经提交过
	count := 0
	if model.DB.Model(&model.DailyInfo{}).Where("uid = ? and save_date = ?", service.Uid, time.Now().Format("2006-01-02")).Count(&count); count > 0 {
		return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
	}

	// 记录用户当日信息
	if err := model.DB.Create(&dailyInfo).Error; err != nil {
		return serializer.ParamErr("上传失败", err)
	}

	return serializer.BuildSuccessSave()
}
