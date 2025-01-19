package repositories

import "sync"

var (
	repositoryInit			sync.Once
	repositoriInstance		*Repository
)

type Repository struct {
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoriInstance = &Repository{
			User:	newUserRepository(),
		}
	})
	return repositoriInstance
}
