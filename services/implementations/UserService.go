package implementations

import "learn/services/interfaces"

type UserService struct {
}

func NewUserService() interfaces.IUserService {
	return &UserService{}
}

func (s *UserService) GetUsernameByID(int) string {
	return "Nimit Didwania"
}
