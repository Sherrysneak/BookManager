// services/borrowing_service.go

package services

import (
	"errors"
	"library/models"
	"library/repositories"
)

type BorrowingService struct {
	borrowingRepo *repositories.BorrowingRepository
	bookRepo      *repositories.BookRepository
}

func NewBorrowingService(borrowingRepo *repositories.BorrowingRepository, bookRepo *repositories.BookRepository) *BorrowingService {
	return &BorrowingService{
		borrowingRepo: borrowingRepo,
		bookRepo:      bookRepo,
	}
}

// BorrowBook 让用户借阅图书
func (s *BorrowingService) BorrowBook(userID uint, bookID uint) error {
	// 检查图书库存
	book, err := s.bookRepo.FindByID(bookID)
	if err != nil {
		return err
	}
	if book.Quantity <= 0 {
		return errors.New("the book is out of stock")
	}

	// 创建新的借阅记录
	borrowing := &models.Borrowing{
		UserID: userID,
		BookID: bookID,
	}
	if err := s.borrowingRepo.Create(borrowing); err != nil {
		return err
	}

	// 更新图书数量
	book.Quantity -= 1
	return s.bookRepo.UpdateBook(book)
}

// ReturnBook 让用户归还图书
func (s *BorrowingService) ReturnBook(userID uint, bookID uint) error {
	// 查找用户借阅的图书记录
	borrowings, err := s.borrowingRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	// 确认用户是否已借阅该图书
	var borrowingToReturn *models.Borrowing
	for _, borrowing := range borrowings {
		if borrowing.BookID == bookID {
			borrowingToReturn = borrowing
			break
		}
	}
	if borrowingToReturn == nil {
		return errors.New("no borrowing record found for this book")
	}

	// 删除借阅记录
	if err := s.borrowingRepo.Delete(borrowingToReturn.ID); err != nil {
		return err
	}

	// 更新图书数量
	book, err := s.bookRepo.FindByID(bookID)
	if err != nil {
		return err
	}
	book.Quantity += 1
	return s.bookRepo.UpdateBook(book)
}

// GetUserBorrowings 查询用户借阅的所有图书
func (s *BorrowingService) GetUserBorrowings(userID uint) ([]*models.Borrowing, error) {
	return s.borrowingRepo.GetByUserID(userID)
}
