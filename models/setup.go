package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := "host=localhost user=postgres password=bismillah dbname=spring port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}