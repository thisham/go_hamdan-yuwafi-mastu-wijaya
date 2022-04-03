package database

import (
	"middleware-api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(configs.SetupDatabase()), &gorm.Config{})

	if err != nil {
		panic("connection to database error!")
	}
	return db
}
