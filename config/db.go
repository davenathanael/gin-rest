package config

import (
	"web/models"

	"github.com/jinzhu/gorm"
)

var dbURL = "user:password@tcp(localhost:9306)/gotest?parseTime=true"

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic("Failed to connect to database, " + err.Error())
	}

	db.AutoMigrate(models.Person{})

	return db
}
