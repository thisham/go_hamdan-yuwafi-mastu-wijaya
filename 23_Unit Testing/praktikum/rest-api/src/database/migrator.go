package database

import (
	bookModel "middleware-api/src/models/book"
	userModel "middleware-api/src/models/user"
)

func Migrate() {
	db := Connect()
	db.AutoMigrate(&userModel.User{}, &bookModel.Book{})
}

func Demigrate() {
	db := Connect()
	db.Migrator().DropTable(&userModel.User{}, &bookModel.Book{})
}
