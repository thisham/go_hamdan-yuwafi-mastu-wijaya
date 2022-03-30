package user

import (
	"gorm.io/gorm"
)

// factory
func NewRepositoryMySQL(db *gorm.DB) UserModel {
	return &UserRepository{
		DB: db,
	}
}
