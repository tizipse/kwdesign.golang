package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebProjectPicture struct {
	Id               int             `gorm:"column:id;primary_key"`
	ClassificationId string          `gorm:"column:classification_id"`
	ProjectId        string          `gorm:"column:project_id"`
	Url              string          `gorm:"column:url"`
	CreatedAt        carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt        carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt  `gorm:"column:deleted_at"`
}
