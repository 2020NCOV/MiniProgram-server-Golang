package serializer

import "ncov_go/model"

// User 用户序列化器
type Info struct {
	IsReturnSchool            string `json:"is_return_school"`
	ReturnTime                string `json:"return_time"`
	ReturnDormNum             string `json:"return_dorm_num"`
	ReturnTrafficInfo         string `json:"return_traffic_info"`
	CurrentHealthValue        string `json:"current_health_value"`
	CurrentContagionRiskValue string `json:"current_contagion_risk_value"`
	ReturnDistrictValue       string `json:"return_district_value"`
	CurrentDistrictValue      string `json:"current_district_value"`
	CurrentTemperature        string `json:"current_temperature"`
	PsyStatus                 string `json:"psy_status"`
	PsyDemand                 string `json:"psy_demand"`
	Remarks                   string `json:"remarks"`
	PsyKnowledge              string `json:"psy_knowledge"`
	PlanCompanyDate           string `json:"plan_company_date"`
}

// BuildInfo 序列化status
func BuildInfo(info model.DailyInfo) Info {
	return Info{
		IsReturnSchool:            info.IsReturnSchool,
		ReturnTime:                info.ReturnTime,
		ReturnDormNum:             info.ReturnDormNum,
		ReturnTrafficInfo:         info.ReturnTrafficInfo,
		CurrentHealthValue:        info.CurrentHealthValue,
		CurrentContagionRiskValue: info.CurrentContagionRiskValue,
		ReturnDistrictValue:       info.ReturnDistrictValue,
		CurrentDistrictValue:      info.CurrentDistrictValue,
		CurrentTemperature:        info.CurrentTemperature,
		PsyStatus:                 info.PsyStatus,
		PsyDemand:                 info.PsyDemand,
		Remarks:                   info.Remarks,
		PsyKnowledge:              info.PsyKnowledge,
		PlanCompanyDate:           info.PlanCompanyDate,
	}
}

// BuildSuccessSave 相应上传成功
func BuildSuccessSave() Response {
	return Response{
		Msg: "上传成功",
	}
}

// BuildLastDataResponse 序列化响应
func BuildLastDataResponse(info model.DailyInfo) Response {
	return Response{
		Data: BuildInfo(info),
	}
}
