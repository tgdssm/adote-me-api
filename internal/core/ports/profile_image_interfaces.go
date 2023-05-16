package ports

import "api/internal/core/domain"

type ProfileImageRepository interface {
	Create(profileImage *domain.ProfileImage) (*domain.ProfileImage, error)
	Update(profileImage *domain.ProfileImage) (*domain.ProfileImage, error)
}

type ProfileImageUseCase interface {
	Create(profileImage *domain.ProfileImage) (*domain.ProfileImage, error)
	Update(profileImage *domain.ProfileImage) (*domain.ProfileImage, error)
}
