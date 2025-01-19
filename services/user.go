package services

import (
	"go-crud/repositories"
	"go-crud/types"
)

type UserService struct {
	userRepository		*repositories.UserRepository
}

func NewUserService(userRep *repositories.UserRepository) *UserService {
	return &UserService {
		userRepository:		userRep,
	}
}

func (u *UserService) CreateUser(newUser *types.User) error {
	return u.userRepository.CreateUser(newUser)
}

func (u *UserService) GetUsers() []*types.User {
	return u.userRepository.GetUsers()
}

func (u *UserService) GetUserByName(name string) (*types.User, error) {
	user, err := u.userRepository.GetUser(name)
	return user, err
}

func (u *UserService) UpdateUser(name, password, updatePassword string) error {
	return u.userRepository.UpdateUser(name, password, updatePassword)
}

func (u *UserService) DeleteUser(name, password string) error {
	return u.userRepository.DeleteUser(name, password)
}
