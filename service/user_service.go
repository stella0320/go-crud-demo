package service

import (
	"go-crud-demo/model"
)

type UserService interface {
	GetAllUsers() []model.User
	CreateNewUser(name string) model.User
	GetUserById(id int) (model.User, error)
}
type userServiceImpl struct{}

var defaultUsers []model.User = []model.User{
	{Id: 1, Name: "John"},
	{Id: 2, Name: "Jane"},
}

func (s *userServiceImpl) GetAllUsers() []model.User {
	return defaultUsers
}

func (s *userServiceImpl) CreateNewUser(name string) model.User {
	newUser := model.User{
		Id:   len(defaultUsers) + 1,
		Name: name,
	}

	defaultUsers = append(defaultUsers, newUser)
	return newUser
}

func (s *userServiceImpl) GetUserById(id int) (model.User, error) {

	var result model.User
	var found = false
	for _, u := range defaultUsers {
		if u.Id == id {
			result = u
			found = true
			break
		}
	}

	if !found {
		return model.User{}, &model.AppError{
			Code:    404,
			Message: "user not found",
		}
	}

	return result, nil
}

func NewUserService() UserService {
	return &userServiceImpl{}
}
