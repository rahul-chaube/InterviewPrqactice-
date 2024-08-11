package user

import "RestAPIDemo/user/model"

type UserInterface interface {
	AddUser(user model.User) error
	GetUser(string) (model.User, error)
	GetUserList() []model.User
}

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) AddUser(user model.User) error {

	return nil
}

func (u *userService) GetUser(id string) (model.User, error) {

	return model.User{}, nil
}
