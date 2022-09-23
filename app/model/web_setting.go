package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebSetting = "web_setting"

type WebSetting struct {
	Id        int             `gorm:"column:id;primary_key"`
	Type      string          `gorm:"column:type"`
	Label     string          `gorm:"column:label"`
	Key       string          `gorm:"column:key"`
	Val       string          `gorm:"column:val"`
	Required  int8            `gorm:"column:required"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const (
	WebSettingTypeOfInput    = "input"
	WebSettingTypeOfTextarea = "textarea"
	WebSettingTypeOfEnable   = "enable"
	WebSettingTypeOfUrl      = "url"
	WebSettingTypeOfEmail    = "email"

	WebSettingRequiredOfYes = 1
	WebSettingRequiredOfNo  = 2
)
