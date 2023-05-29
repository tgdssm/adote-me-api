package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type PetPhotoUseCase struct {
	repo ports.PetPhotoUseCase
}

func NewPetPhotoUseCase(repo ports.PetPhotoUseCase) *PetPhotoUseCase {
	return &PetPhotoUseCase{
		repo: repo,
	}
}

func (u PetPhotoUseCase) Create(petPhoto *domain.PetPhoto) (*domain.PetPhoto, error) {
	return u.repo.Create(petPhoto)
}
