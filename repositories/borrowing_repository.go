// repository/borrowing_repository.go

package repositories

import (
	"gorm.io/gorm"
	"library/models"
)

type BorrowingRepository struct {
	DB *gorm.DB
}

func NewBorrowingRepository(db *gorm.DB) *BorrowingRepository {
	return &BorrowingRepository{DB: db}
}

func (r *BorrowingRepository) Create(borrowing *models.Borrowing) error {
	return r.DB.Create(borrowing).Error
}

func (r *BorrowingRepository) GetByID(id uint) (*models.Borrowing, error) {
	var borrowing models.Borrowing
	err := r.DB.First(&borrowing, id).Error
	return &borrowing, err
}

func (r *BorrowingRepository) Update(borrowing *models.Borrowing) error {
	return r.DB.Save(borrowing).Error
}

func (r *BorrowingRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Borrowing{}, id).Error
}

func (r *BorrowingRepository) GetByUserID(userID uint) ([]*models.Borrowing, error) {
	var borrowings []*models.Borrowing
	err := r.DB.Where("user_id = ?", userID).Find(&borrowings).Error
	return borrowings, err
}
