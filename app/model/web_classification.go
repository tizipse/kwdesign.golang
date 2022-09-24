package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebClassification = "web_classification"

type WebClassification struct {
	Id          string          `gorm:"column:id;primary_key"`
	Name        string          `gorm:"column:name"`
	Title       string          `gorm:"column:title"`
	Keyword     string          `gorm:"column:keyword"`
	Description string          `gorm:"column:description"`
	Order       int8            `gorm:"column:order"`
	IsEnable    int8            `gorm:"column:is_enable"`
	CreatedAt   carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt   carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"column:deleted_at"`
}
