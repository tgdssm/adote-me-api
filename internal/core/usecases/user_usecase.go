package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type userUseCase struct {
	repo ports.UserRepository
}

func NewUserUseCase(repo ports.UserRepository) *userUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u userUseCase) Create(user *domain.User) (*domain.User, error) {
	return u.repo.Create(user)
}

func (u userUseCase) List(queryParameter string) ([]domain.User, error) {
	return u.repo.List(queryParameter)
}

func (u userUseCase) Get(id int) (*domain.User, error) {
	return u.repo.Get(id)
}
