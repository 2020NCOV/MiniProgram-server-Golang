package model

import "github.com/jinzhu/gorm"

type DailyInfo struct {
	gorm.Model
	IsReturnSchool            string
	CurrentHealthValue        string
	CurrentContagionRiskValue string
	ReturnDistrictValue       string
	CurrentDistrictValue      string
	CurrentTemperature        string
	Remarks                   string
	PsyStatus                 string
	PsyDemand                 string
	PsyKnowledge              string
	PlanCompanyDate           string
	ReturnTime                string
	ReturnDormNum             string
	ReturnTrafficInfo         string
	Uid                       string
	SaveDate                  string
}
