package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Due         time.Time `gorm:"index"`
	PlanID      uint
	Plan        Plan
}
