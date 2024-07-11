package implementations

import (
	"learn/models"
	"learn/services/interfaces"
)

type UserService struct {
}

func NewUserService() interfaces.IUserService {
	return &UserService{}
}

func (s *UserService) GetUser(id int) models.User {
	return models.User{}
}

func (s *UserService) GetUsers(request models.GetAllItemsRequest) []models.User {
	panic("not implemented")
}

func (s *UserService) CreateUser(user models.User) {

}

func (s *UserService) DeleteUser(id int) {

}

func (s *UserService) ValidateCredentials(user models.User) {

}
