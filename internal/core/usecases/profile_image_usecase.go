package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type ProfileImageUseCase struct {
	repo ports.ProfileImageRepository
}

func NewProfileImageUseCase(repo ports.ProfileImageUseCase) *ProfileImageUseCase {
	return &ProfileImageUseCase{
		repo: repo,
	}
}

func (p ProfileImageUseCase) Create(profileImage *domain.ProfileImage) (*domain.ProfileImage, error) {
	return p.repo.Create(profileImage)
}

func (p ProfileImageUseCase) Update(profileImage *domain.ProfileImage) (*domain.ProfileImage, error) {
	return p.repo.Update(profileImage)
}
