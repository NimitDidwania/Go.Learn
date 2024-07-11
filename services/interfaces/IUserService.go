package interfaces

import "learn/models"

type IUserService interface {
	// Get by ID
	GetUser(int) models.User
	// Get All Elements with search and pagination support
	GetUsers(models.GetAllItemsRequest) []models.User
	// Create User
	CreateUser(models.User)
	// Delete User
	DeleteUser(int)
	// Validate user credentials
	ValidateCredentials(models.User)
}
