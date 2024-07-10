package interfaces

import "learn/models"

type IUserService interface {
	GetUser(int) models.User
	GetUsers()
}
