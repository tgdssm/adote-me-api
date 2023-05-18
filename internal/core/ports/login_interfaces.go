package ports

import "api/internal/core/domain"

type LoginRepository interface {
	GetByEmail(email string) (*domain.User, error)
}

type LoginUseCase interface {
	GetByEmail(email string) (*domain.User, error)
}
