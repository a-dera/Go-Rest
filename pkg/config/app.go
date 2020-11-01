package config

import (
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() {
	// Please define your user name and password for mysql.
	d, err := gorm.Open("mysql", "root:@/gorest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	database = d
}

func GetDB() *gorm.DB {
	return database
}
