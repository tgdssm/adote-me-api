package ports

import "api/internal/core/domain"

type PetRepository interface {
	Create(pet *domain.Pet) (*domain.Pet, error)
	List(queryParameter string) ([]domain.Pet, error)
	ListByUser(userID int) ([]domain.Pet, error)
	Get(id int) (*domain.Pet, error)
	Delete(id int) error
}

type PetUseCase interface {
	Create(pet *domain.Pet) (*domain.Pet, error)
	List(queryParameter string) ([]domain.Pet, error)
	ListByUser(userID int) ([]domain.Pet, error)
	Get(id int) (*domain.Pet, error)
	Delete(id int) error
}
