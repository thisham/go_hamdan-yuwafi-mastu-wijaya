package database

import (
	userModel "gorm-mvc-api/src/models/user"
)

func Migrate() {
	db := Connect()
	db.AutoMigrate(&userModel.User{})
}
