package models

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	gorm.Model
	Name          string
	StartDate     time.Time
	EndDate       time.Time
	NumberOfWeeks int
	Tasks         *[]Task
	Todos         *[]Todo
}
