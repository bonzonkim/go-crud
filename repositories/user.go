package repositories

import (
	"fmt"
	"go-crud/types"
)

type UserRepository struct {
	userMap []*types.User
	user	*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository {
		userMap:	[]*types.User{},
	}
}

func (u *UserRepository) CreateUser(user *types.User) error {
	u.userMap = append(u.userMap, user)
	fmt.Println(u.userMap)
	return nil
}

func (u *UserRepository) GetUsers() []*types.User {
	return u.userMap
}

func (u *UserRepository) GetUser(name string) (*types.User, error) {
	isExisit := false
	for _, user := range u.userMap {
		if user.Name == name {
			isExisit = true
			continue
		}
	}

	if !isExisit {
		return nil, types.Errorf(types.NotFoundUser, nil)
	}
	return u.user, nil
}

func (u *UserRepository) UpdateUser(name string, password string, updatePassword string) error {
	isExisit := false

	for _, user := range u.userMap {
		if user.Name == name && user.Password == password {
			user.Password = updatePassword	
			isExisit = true
			continue
		}
	}
	if !isExisit {
		return types.Errorf(types.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) DeleteUser(name string, password string) error {
	isExisit := false

	for i, user := range u.userMap {
		if user.Name == name && user.Password == password {
			u.userMap = append(u.userMap[:i], u.userMap[i+1:]...)
			isExisit = true
			continue
		}
	}
	if !isExisit {
		return types.Errorf(types.NotFoundUser, nil)
	} else {
		return nil
	}
}
