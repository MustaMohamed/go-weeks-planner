package database

import (
	"github.com/mustamohamed/weekplanner/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	RegisterConnection()
}

func RegisterConnection() {
	var err error
	Connection, err = gorm.Open(sqlite.Open("WeekPlanner"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Connection.AutoMigrate(&models.Plan{})
	Connection.AutoMigrate(&models.Task{})
	Connection.AutoMigrate(&models.Todo{})
}

func GetConnection() *gorm.DB {
	return Connection
}
