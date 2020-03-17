package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"log"

	"github.com/gin-gonic/gin"
)

// 管理获取表单数据服务
type GetLastDataService struct {
	UID   int    `form:"uid" json:"uid"`
	Token string `form:"token" json:"token"`
}

// 获取上次提交的数据
func (service *GetLastDataService) GetLastData(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.UID, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//获取用户所属的机构
	var templateCode string
	err := model.DB.QueryRow(`select o.template_code 
		from wx_mp_bind_info u
		left join organization o
		on u.org_id = o.id
		where u.wx_uid = ? and u.isbind = 1`, service.UID).
		Scan(&templateCode)
	if err != nil {
		return serializer.Err(10006, "获取用户绑定信息失败", nil)
	}

	// 获取表单信息
	var lastData model.Record
	queryStr := `select is_return_school,IFNULL(remarks,""),IFNULL(return_dorm_num,""),IFNULL(return_time,""),IFNULL(return_traffic_info,""),
	IFNULL(current_health_value,0),IFNULL(current_contagion_risk_value,0),IFNULL(return_district_value,0),IFNULL(current_district_value,0),
	IFNULL(current_temperature,0),IFNULL(psy_status,0),IFNULL(psy_demand,0),IFNULL(psy_knowledge,0),IFNULL(plan_company_date,"") 
	from ` + "report_record_" + templateCode + " where wxuid = ? order by time desc limit 1"
	err = model.DB.QueryRow(queryStr, service.UID).Scan(&lastData.IsReturnSchool, &lastData.Remarks, &lastData.ReturnDormNum,
		&lastData.ReturnTime, &lastData.ReturnTrafficInfo, &lastData.CurrentHealthValue, &lastData.CurrentContagionRiskValue,
		&lastData.ReturnDistrictValue, &lastData.CurrentDistrictValue, &lastData.CurrentTemperature, &lastData.PsyStatus,
		&lastData.PsyDemand, &lastData.PsyKnowledge, &lastData.PlanCompanyDate)
	if err == nil {
		lastData.ReturnDistrictPath = service.getDistrictPath(lastData.ReturnDistrictValue)
		lastData.CurrentDistrictPath = service.getDistrictPath(lastData.CurrentDistrictValue)
		return serializer.BuildLastDataResponse(false, lastData)
	}
	// 出现错误时返回空值
	log.Println(err)
	return serializer.BuildLastDataResponse(true, lastData)
}

// getDistrictPath 获取行政区信息
func (service *GetLastDataService) getDistrictPath(cityCode int) string {
	var dis model.District
	var pathStr string
	err := model.DB.QueryRow("select name,level_id,parent_id from com_district where value = ?", cityCode).
		Scan(&dis.Name, &dis.LevelID, &dis.ParentID)
	if err == nil {
		pathStr = dis.Name
		if dis.LevelID != 1 {
			err = model.DB.QueryRow("select name from com_district where value = ?", dis.ParentID).
				Scan(&dis.Name)
			if err == nil {
				pathStr = dis.Name + pathStr
			}
		}
	}
	return pathStr
}
