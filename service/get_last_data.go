package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetLastDataService 管理获取表单数据服务
type GetLastDataService struct {
	UID   int    `form:"uid" json:"uid"`
	Token string `form:"token" json:"token"`
}

// GetLastData 获取上次提交的数据
func (service *GetLastDataService) GetLastData(c *gin.Context) serializer.Response {

	if !model.CheckToken(strconv.Itoa(service.UID), service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//获取用户所属的机构
	var templateCode string
	err := model.DB2.QueryRow(`select o.template_code 
		from wx_mp_bind_info u
		left join organization o
		on u.org_id = o.id
		where u.wx_uid = ? and u.isbind = 1`, service.UID).
		Scan(&templateCode)
	if err != nil {
		return serializer.Err(10006, "获取用户绑定信息失败", nil)
	}

	// 获取表单信息
	var lastData model.DailyInfo
	queryStr := `select is_return_school,remarks,return_dorm_num,return_time,return_traffic_info,current_health_value,
	current_contagion_risk_value,return_district_value,current_district_value,current_temperature,psy_status,
	psy_demand,psy_knowledge,plan_company_date from ` +
		"report_record_" + templateCode + " where wxuid = ? order by time desc limit 1"
	model.DB2.QueryRow(queryStr, service.UID).Scan(&lastData.IsReturnSchool, &lastData.Remarks, &lastData.ReturnDormNum,
		&lastData.ReturnTime, &lastData.ReturnTrafficInfo, &lastData.CurrentHealthValue, &lastData.CurrentContagionRiskValue,
		&lastData.ReturnDistrictValue, &lastData.CurrentDistrictValue, &lastData.CurrentTemperature, &lastData.PsyStatus,
		&lastData.PsyDemand, &lastData.PsyKnowledge, &lastData.PlanCompanyDate)
	if err == nil {
		//log.Println(lastData.ReturnDistrictValue + " " + lastData.CurrentDistrictValue)
		lastData.ReturnDistrictPath = service.getDistrictPath(lastData.ReturnDistrictValue)
		lastData.CurrentDistrictPath = service.getDistrictPath(lastData.CurrentDistrictValue)
		return serializer.BuildLastDataResponse(false, lastData)
	}

	return serializer.BuildLastDataResponse(true, lastData)
}

// getDistrictPath 获取行政区信息
func (service *GetLastDataService) getDistrictPath(cityCode string) string {
	var dis model.District
	code, err := strconv.Atoi(cityCode)
	err = model.DB2.QueryRow("select name,level_id,parent_id from com_district where value = ?", code).
		Scan(&dis.Name, &dis.LevelID, &dis.ParentID)
	var pathStr string
	if err == nil {
		pathStr = dis.Name
		if dis.LevelID != 1 {
			err = model.DB2.QueryRow("select name from com_district where value = ?", dis.ParentID).
				Scan(&dis.Name)
			if err == nil {
				pathStr = dis.Name + pathStr
			}
		}
	}
	return pathStr
}
