package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"fmt"
	"log"
)

type ProfileImageMysqlRepository struct {
	db *sql.DB
}

func NewProfileImageMysqlRepository() *ProfileImageMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return &ProfileImageMysqlRepository{
		db: db,
	}
}

func (repo ProfileImageMysqlRepository) Create(profileImage *domain.ProfileImage) (*domain.ProfileImage, error) {
	statement, err := repo.db.Prepare("insert into profile_images(file_name, file_path, user_id) values(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()
	fmt.Println(profileImage)
	insert, err := statement.Exec(profileImage.FileName, profileImage.FilePath, profileImage.UserID)
	if err != nil {
		return nil, err
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	profileImage.ID = uint64(id)

	return profileImage, nil
}

func (repo ProfileImageMysqlRepository) Update(profileImage *domain.ProfileImage) (*domain.ProfileImage, error) {
	statement, err := repo.db.Prepare("update profile_images set file_name = ?, file_path = ? where id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	_, err = statement.Exec(profileImage.FileName, profileImage.FilePath, profileImage.ID)
	if err != nil {
		return nil, err
	}

	return profileImage, nil
}
