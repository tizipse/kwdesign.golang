package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebContact = "web_contact"

type WebContact struct {
	Id        int             `gorm:"column:id;primary_key"`
	City      string          `gorm:"column:city"`
	Address   string          `gorm:"column:address"`
	Telephone string          `gorm:"column:telephone"`
	Order     int8            `gorm:"column:order"`
	IsEnable  int8            `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}
