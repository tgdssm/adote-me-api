package ports

import "api/internal/core/domain"

type PetPhotoRepository interface {
	Create(profileImage *domain.PetPhoto) (*domain.PetPhoto, error)
}

type PetPhotoUseCase interface {
	Create(profileImage *domain.PetPhoto) (*domain.PetPhoto, error)
}
