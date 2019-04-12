package config

import (
	"github.com/jinzhu/gorm"
	"github.com/mahdifr17/phonebook-v2/structs"
)

var gormInstance *gorm.DB

// GetDbInstance is a singleton
func GetDbInstance() *gorm.DB {
	if gormInstance == nil {
		gormInstance = dbInit()
	}

	return gormInstance
}

func dbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "{MYSQL_USERNAME}:{MYSQL_PASSWORD}@tcp(127.0.0.1:3306)/phonebook?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("Database connection failed")
	} else {
		db.AutoMigrate(structs.Record{})
		return db
	}
}
