package ports

import "api/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	List(queryParameter string) ([]domain.User, error)
	Get(id int) (*domain.User, error)
}

type UserUseCase interface {
	Create(user *domain.User) (*domain.User, error)
	List(queryParameter string) ([]domain.User, error)
	Get(id int) (*domain.User, error)
}
