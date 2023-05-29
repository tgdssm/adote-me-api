package usecases

import (
	"api/internal/core/domain"
	"api/internal/core/ports"
)

type PetUseCase struct {
	repo ports.PetRepository
}

func NewPetUseCase(repo ports.PetRepository) *PetUseCase {
	return &PetUseCase{
		repo: repo,
	}
}

func (u PetUseCase) Create(pet *domain.Pet) (*domain.Pet, error) {
	return u.repo.Create(pet)
}

func (u PetUseCase) List(queryParameter string) ([]domain.Pet, error) {
	return u.repo.List(queryParameter)
}

func (u PetUseCase) Get(id int) (*domain.Pet, error) {
	return u.repo.Get(id)
}
