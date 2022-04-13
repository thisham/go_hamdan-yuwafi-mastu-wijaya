package auth

import "gorm.io/gorm"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginModel interface {
	AttemptLogin(data Login) (string, error)
}

type LoginRepository struct {
	DB *gorm.DB
}
