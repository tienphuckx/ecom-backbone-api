package service

import (
	"github.com/tienphuckx/ecom-backbone-api.git/internal/repo"
)

// UserService is a service that interacts with the UserRepository.
type UserService struct {
	UserRepository *repo.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService() *UserService {
	// The service initializes the repository
	return &UserService{
		UserRepository: repo.NewUserRepository(),
	}
}

// GetUser retrieves user information by ID.
func (us *UserService) GetUser(userID string) string {
	// Pass the userID to the repository
	return us.UserRepository.GetUserInfo(userID)
}

// GetUserByEmail retrieves user information by email.
func (us *UserService) GetUserByEmail(userEmail string) string {
	return us.UserRepository.GetUserInfoByEmail(userEmail)
}
