package repositories

import (
	"gorm.io/gorm"
	"library/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(userID uint) error {
	return r.DB.Delete(&models.User{}, userID).Error
}

func (r *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
