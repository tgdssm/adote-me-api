package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type LoginUseCase struct {
	repo ports.LoginRepository
}

func NewLoginUseCase(repo ports.LoginRepository) *LoginUseCase {
	return &LoginUseCase{
		repo: repo,
	}
}

func (login LoginUseCase) GetByEmail(email string) (*domain.User, error) {
	return login.repo.GetByEmail(email)
}
