package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// create uuid on create triggered
func (base *User) BeforeCreate(tx *gorm.DB) error {
	uuid, _ := uuid.NewRandom()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
