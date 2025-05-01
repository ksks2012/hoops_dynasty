package global

import (
	"github.com/hoops_dynasty/pkg/logger"
	"github.com/hoops_dynasty/pkg/setting"
)

var (
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	SaveLoadSetting *setting.SaveLoadSettingS
	Logger          *logger.Logger
	// Limits
	BasicSetting    *setting.BasicSettingS
	DefenseSetting  *setting.DefenseSettingS
	MentalSetting   *setting.MentalSettingS
	OffenseSetting  *setting.OffenseSettingS
	PhysicalSetting *setting.PhysicalSettingS
)
