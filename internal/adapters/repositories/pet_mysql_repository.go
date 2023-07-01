package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PetMysqlRepository struct {
	db *sql.DB
}

func NewPetMysqlRepository() *PetMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return &PetMysqlRepository{
		db: db,
	}
}

func (repo PetMysqlRepository) Create(pet *domain.Pet) (*domain.Pet, error) {
	statement, err := repo.db.Prepare("insert into pets(pet_name, age, weight, requirements, user_id) values(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	insert, err := statement.Exec(pet.Name, pet.Age, pet.Weight, pet.Requirements, pet.User.ID)
	if err != nil {
		return nil, err
	}

	lastId, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	pet.ID = uint64(lastId)

	return pet, nil
}

func (repo PetMysqlRepository) List(queryParameter string) ([]domain.Pet, error) {
	var rows *sql.Rows
	var err error
	if queryParameter == "" {
		rows, err = repo.db.Query("select p.id, p.pet_name, p.age, p.weight, p.requirements, p.created_at, pi.id, pi.file_name, pi.file_path, u.id, u.username, u.email, u.cellphone, u.created_at from pets p inner join pet_images pi on p.id = pi.pet_id inner join users u on p.user_id = u.id")

	} else {
		queryParameter = fmt.Sprintf("%%%s%%", queryParameter)
		rows, err = repo.db.Query("select p.id, p.pet_name, p.age, p.weight, p.requirements, p.created_at, pi.id, pi.file_name, pi.file_path, u.id, u.username, u.email, u.cellphone, u.created_at from pets p inner join pet_images pi on p.id = pi.pet_id inner join users u on p.user_id = u.id and p.pet_name like ?", queryParameter)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pets := []domain.Pet{}
	var pet domain.Pet

	var petPhoto domain.PetPhoto
	for rows.Next() {
		if err = rows.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Weight, &pet.Requirements, &pet.CreatedAt, &petPhoto.ID, &petPhoto.FileName, &petPhoto.FilePath, &pet.User.ID, &pet.User.Name, &pet.User.Email, &pet.User.Cellphone, &pet.User.CreatedAt); err != nil {
			return nil, err
		}

		var contains bool
		var index int
		for i, p := range pets {
			if p.ID == pet.ID {
				contains = true
				index = i
				break
			}
		}

		if contains {
			pets[index].Photos = append(pets[index].Photos, petPhoto)
		} else {
			// Limpar o slice antes de adicionar uma nova foto de um outro pet
			pet.Photos = []domain.PetPhoto{petPhoto}
			pets = append(pets, pet)
		}
	}

	return pets, nil
}

func (repo PetMysqlRepository) Get(id int) (*domain.Pet, error) {
	var pet domain.Pet

	row, err := repo.db.Query("select * from pets p inner join pet_images pi on p.id = pi.pet_id where p.id = ?", id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var petPhoto domain.PetPhoto

	for row.Next() {
		if err = row.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Weight, &pet.Requirements, &pet.CreatedAt, &petPhoto.ID, &petPhoto.FileName, &petPhoto.FilePath, &petPhoto.PetID, &pet.ID, &pet.Name, &pet.Age, &pet.Weight, &pet.Requirements, &petPhoto.ID, &petPhoto.FileName, &petPhoto.FilePath, &petPhoto.PetID, &pet.User.ID, &pet.User.Name, &pet.User.Email, &pet.User.Cellphone, &pet.User.CreatedAt); err != nil {
			return nil, err
		}
		pet.Photos = append(pet.Photos, petPhoto)
	}

	return &pet, nil
}
