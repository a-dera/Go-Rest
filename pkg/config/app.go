package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	database * gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	d, err := gorm.Open("mysql", "root:@/gorest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	database = d
}

func GetDB() *gorm.DB {
	return database
}
