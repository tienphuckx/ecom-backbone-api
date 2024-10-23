package repo

// UserRepository handles the data access for user information.
type UserRepository struct{}

// NewUserRepository creates a new UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetUserInfo retrieves user information from the data source.
func (ur *UserRepository) GetUserInfo(userID string) string {
	return "User Information for ID: " + userID
}

// GetUserInfoByEmail retrieves user information from the data source.
func (ur *UserRepository) GetUserInfoByEmail(userEmail string) string {
	return "User Information for Email: " + userEmail
}
