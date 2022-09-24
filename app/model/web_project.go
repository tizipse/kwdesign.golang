package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebProject struct {
	Id               string          `gorm:"column:id;primary_key"`
	ClassificationId string          `gorm:"column:classification_id"`
	Theme            string          `gorm:"column:theme"`
	Name             string          `gorm:"column:name"`
	Address          string          `gorm:"column:address"`
	Picture          string          `gorm:"column:picture"`
	Title            string          `gorm:"column:title"`
	Keyword          string          `gorm:"column:keyword"`
	Description      string          `gorm:"column:description"`
	Html             string          `gorm:"column:html"`
	IsEnable         int8            `gorm:"column:is_enable"`
	DatedAt          carbon.Date     `gorm:"column:dated_at"`
	CreatedAt        carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt        carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt  `gorm:"column:deleted_at"`

	Classification *WebClassification  `gorm:"foreignKey:Id;references:ClassificationId"`
	Pictures       []WebProjectPicture `gorm:"foreignKey:ProjectId;references:Id"`
}
