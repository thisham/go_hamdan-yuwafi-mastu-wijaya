package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel interface {
	GetAllUsers() ([]User, error)
	GetUser(id string) (User, error)
}

type UserRepository struct {
	DB *gorm.DB
}
