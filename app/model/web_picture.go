package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebPicture = "web_picture"

type WebPicture struct {
	Id        int             `gorm:"column:id;primary_key"`
	Label     string          `gorm:"column:label"`
	Key       string          `gorm:"column:key"`
	Val       string          `gorm:"column:val"`
	Required  int8            `gorm:"column:required"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const (
	WebPictureRequiredOfYes = 1
	WebPictureRequiredOfNo  = 2
)
