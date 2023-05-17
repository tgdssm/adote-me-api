package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type UserUseCase struct {
	repo ports.UserRepository
}

func NewUserUseCase(repo ports.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u UserUseCase) Create(user *domain.User) (*domain.User, error) {
	return u.repo.Create(user)
}

func (u UserUseCase) List(queryParameter string) ([]domain.User, error) {
	return u.repo.List(queryParameter)
}

func (u UserUseCase) Get(id int) (*domain.User, error) {
	return u.repo.Get(id)
}

func (u UserUseCase) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u UserUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}
