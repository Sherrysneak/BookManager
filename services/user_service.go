package services

import (
	"errors"
	"library/models"
	"library/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) CreateUser(user *models.User) error {
	// Add any additional business logic here (e.g., validating user data)
	return u.userRepo.CreateUser(user)
}

func (u *UserService) GetUserByID(id uint) (*models.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *UserService) UpdateUser(user *models.User) error {
	// Add additional checks (e.g., checking if the user exists)
	return u.userRepo.UpdateUser(user)
}

func (u *UserService) DeleteUser(id uint) error {
	// Ensure the user exists before deleting
	_, err := u.userRepo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}
	return u.userRepo.DeleteUser(id)
}
