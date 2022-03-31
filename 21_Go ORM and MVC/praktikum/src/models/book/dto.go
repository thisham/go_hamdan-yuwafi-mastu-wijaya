package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        uuid.UUID `json:"ID" gorm:"type:string;size:36;primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	ISBN      string    `json:"isbn"`
	Language  string    `json:"language"`
}

type BookModel interface {
}

type BookRepository struct {
	DB *gorm.DB
}
