package factories

import (
	"gorm-mvc-api/src/models/user"

	"gorm.io/gorm"
)

func UserFactory(db *gorm.DB) user.UserModel {
	return &user.UserRepository{
		DB: db,
	}
}
