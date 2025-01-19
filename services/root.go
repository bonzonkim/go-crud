package services

import (
	"go-crud/repositories"
	"sync"
)

var (
	serviceInit			sync.Once
	serviceInstance		*Service
)

type Service struct {
	repository		*repositories.Repository
	User			*UserService
}

func NewService(rep *repositories.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance := &Service{
			repository:		rep,
		}
		serviceInstance.User = NewUserService(rep.User)
	})
	return serviceInstance
}
