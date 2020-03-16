package serializer

import "Miniprogram-server-Golang/model"

// 用户序列化器
type Record struct {
	IsReturnSchool            int    `json:"is_return_school"`
	CurrentHealthValue        int    `json:"current_health_value"`
	CurrentContagionRiskValue int    `json:"current_contagion_risk_value"`
	ReturnDistrictValue       int    `json:"return_district_value"`
	CurrentDistrictValue      int    `json:"current_district_value"`
	CurrentTemperature        int    `json:"current_temperature"`
	PsyStatus                 int    `json:"psy_status"`
	PsyDemand                 int    `json:"psy_demand"`
	PsyKnowledge              int    `json:"psy_knowledge"`
	ReturnTime                string `json:"return_time"`
	ReturnDormNum             string `json:"return_dorm_num"`
	ReturnTrafficInfo         string `json:"return_traffic_info"`
	Remarks                   string `json:"remarks"`
	PlanCompanyDate           string `json:"plan_company_date"`
	ReturnDistrictPath        string `json:"return_district_path"`
	CurrentDistrictPath       string `json:"current_district_path"`
}

// 生成返回消息体
type ResponseData struct {
	Errcode int         `json:"errcode"`
	IsEmpty int         `json:"isEmpty"`
	Data    interface{} `json:"data"`
}

//  序列化report
func BuildRecord(record model.Record) Record {
	return Record{
		IsReturnSchool:            record.IsReturnSchool,
		ReturnTime:                record.ReturnTime,
		ReturnDormNum:             record.ReturnDormNum,
		ReturnTrafficInfo:         record.ReturnTrafficInfo,
		CurrentHealthValue:        record.CurrentHealthValue,
		CurrentContagionRiskValue: record.CurrentContagionRiskValue,
		ReturnDistrictValue:       record.ReturnDistrictValue,
		CurrentDistrictValue:      record.CurrentDistrictValue,
		CurrentTemperature:        record.CurrentTemperature,
		PsyStatus:                 record.PsyStatus,
		PsyDemand:                 record.PsyDemand,
		Remarks:                   record.Remarks,
		PsyKnowledge:              record.PsyKnowledge,
		PlanCompanyDate:           record.PlanCompanyDate,
		ReturnDistrictPath:        record.ReturnDistrictPath,
		CurrentDistrictPath:       record.CurrentDistrictPath,
	}
}

//  相应上传成功
func BuildSuccessSave() Response {
	return Response{
		Msg: "上传成功",
	}
}

//  序列化响应
func BuildLastDataResponse(isEmpty bool, record model.Record) Response {
	if isEmpty {
		return Response{
			Data: ResponseData{0, 1, nil},
		}
	}
	return Response{
		Data: ResponseData{0, 0, BuildRecord(record)},
	}
}
