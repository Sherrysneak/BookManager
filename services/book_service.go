package services

import (
	"library/models"
	"library/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(bookRepo *repositories.BookRepository) *BookService {
	return &BookService{repo: bookRepo}
}
func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(bookID uint) error {
	return s.repo.DeleteBook(bookID)
}

func (s *BookService) GetBookByID(bookID uint) (*models.Book, error) {
	return s.repo.FindByID(bookID)
}
