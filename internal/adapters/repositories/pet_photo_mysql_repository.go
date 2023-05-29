package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"log"
)

type PetPhotoMysqlRepository struct {
	db *sql.DB
}

func NewPetPhotoMysqlRepository() *PetPhotoMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return &PetPhotoMysqlRepository{
		db: db,
	}
}

func (repo PetPhotoMysqlRepository) Create(petPhoto *domain.PetPhoto) (*domain.PetPhoto, error) {
	statement, err := repo.db.Prepare("insert into pet_images(file_name, file_path, pet_id) values(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	insert, err := statement.Exec(petPhoto.FileName, petPhoto.FilePath, petPhoto.PetID)
	if err != nil {
		return nil, err
	}

	lastId, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	petPhoto.ID = uint64(lastId)

	return petPhoto, nil
}
