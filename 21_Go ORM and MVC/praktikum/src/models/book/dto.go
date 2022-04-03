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
	GetAllBooks() ([]Book, error)
	GetBook(id string) (Book, error)
	CreateBook(book Book) error
	UpdateBook(book Book, id string) error
	DeleteBook(id string) error
}

type BookRepository struct {
	DB *gorm.DB
}
