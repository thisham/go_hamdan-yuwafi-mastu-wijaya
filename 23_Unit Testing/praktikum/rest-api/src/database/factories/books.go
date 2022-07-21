package factories

import (
	"middleware-api/src/models/book"

	"gorm.io/gorm"
)

func BookFactory(db *gorm.DB) book.BookModel {
	return &book.BookRepository{
		DB: db,
	}
}
