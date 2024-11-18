package repositories

import (
	"gorm.io/gorm"
	"library/models"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}
func (r *BookRepository) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) UpdateBook(book *models.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepository) DeleteBook(bookID uint) error {
	return r.DB.Delete(&models.Book{}, bookID).Error
}

func (r *BookRepository) FindByID(bookID uint) (*models.Book, error) {
	var book models.Book
	if err := r.DB.First(&book, bookID).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
