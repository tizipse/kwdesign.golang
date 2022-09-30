package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebCategory = "web_category"

type WebCategory struct {
	Id                int             `gorm:"column:id;primary_key"`
	Uri               string          `gorm:"column:uri"`
	Name              string          `gorm:"column:name"`
	Picture           string          `gorm:"column:picture"`
	Title             string          `gorm:"column:title"`
	Keyword           string          `gorm:"column:keyword"`
	Description       string          `gorm:"column:description"`
	Html              string          `gorm:"column:html"`
	IsRequiredPicture int8            `gorm:"column:is_required_picture"`
	IsRequiredHtml    int8            `gorm:"column:is_required_html"`
	IsEnable          int8            `gorm:"column:is_enable"`
	CreatedAt         carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt         carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const (
	WebCategoryOfIsRequiredPictureYes = 1
	WebCategoryOfIsRequiredPictureNo  = 2
	WebCategoryOfIsRequiredHtmlYes    = 1
	WebCategoryOfIsRequiredHtmlNo     = 2
)
