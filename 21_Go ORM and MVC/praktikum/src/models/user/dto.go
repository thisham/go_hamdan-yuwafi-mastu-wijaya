package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserModel interface {
	GetAllUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user User) error
	UpdateUser(user User, id string) error
}

type UserRepository struct {
	DB *gorm.DB
}
