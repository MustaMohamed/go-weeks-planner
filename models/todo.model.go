package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Name   string
	Due    time.Time `gorm:"index"`
	PlanID uint
	Plan   Plan
}
