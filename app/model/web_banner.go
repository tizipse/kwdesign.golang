package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebBanner = "web_banner"

type WebBanner struct {
	Id        int             `gorm:"column:id;primary_key"`
	Client    string          `gorm:"column:client"`
	Theme     string          `gorm:"column:theme"`
	Picture   string          `gorm:"column:picture"`
	Name      string          `gorm:"column:name"`
	Target    string          `gorm:"column:target"`
	Url       string          `gorm:"column:url"`
	IsEnable  int8            `gorm:"column:is_enable"`
	Order     int8            `gorm:"column:order"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const (
	WebBannerClientOfPc     = "PC"
	WebBannerClientOfMobile = "MOBILE"
	WebBannerTargetOfBlank  = "blank"
	WebBannerTargetOfSelf   = "self"
)
