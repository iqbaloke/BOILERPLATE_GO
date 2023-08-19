package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connection() {
	dsn := "sqlserver://SA:yourStrong(!)Password@localhost:1433?database=belajar_go"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild Connect To Database")
	}
	DB = db

}
