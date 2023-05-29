package ports

import "api/internal/core/domain"

type PetRepository interface {
	Create(pet *domain.Pet) (*domain.Pet, error)
	List(queryParameter string) ([]domain.Pet, error)
	Get(id int) (*domain.Pet, error)
}

type PetUseCase interface {
	Create(pet *domain.Pet) (*domain.Pet, error)
	List(queryParameter string) ([]domain.Pet, error)
	Get(id int) (*domain.Pet, error)
}
