package book

func (repo *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	if err := repo.DB.Find(&books).Error; err != nil {
		return []Book{}, err
	}

	return books, nil
}

func (repo *BookRepository) GetBook(id string) (Book, error) {
	var book Book
	if err := repo.DB.Where("ID = ?", id).Find(&book).Error; err != nil {
		return Book{}, err
	}

	return book, nil
}

func (repo *BookRepository) CreateBook(book Book) error {
	if err := repo.DB.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BookRepository) UpdateBook(book Book, id string) error {
	if err := repo.DB.Where("ID = ?", id).Updates(&book).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BookRepository) DeleteBook(id string) error {
	if err := repo.DB.Where("ID = ?", id).Delete(&Book{}).Error; err != nil {
		return err
	}

	return nil
}
