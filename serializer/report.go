package serializer

import "Miniprogram-server-Golang/model"

// Record 用户序列化器
type Record struct {
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

// BuildRecord 序列化report
func BuildRecord(info model.DailyInfo) Record {
	return Record{
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
func BuildLastDataResponse(isEmpty bool, info model.DailyInfo) Response {
	if isEmpty {
		return Response{
			Data: "",
		}
	}
	return Response{
		Data: BuildRecord(info),
	}
}
