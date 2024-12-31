package models

import (
	"gorm.io/gorm"
)

type BasicModel struct {
	gorm.Model
	IsActive bool `gorm:"default:true" json:"isActive"`
}
