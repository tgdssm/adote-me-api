package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"log"
)

type LoginMysqlRepository struct {
	db *sql.DB
}

func NewLoginMysqlRepository() *LoginMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return &LoginMysqlRepository{
		db: db,
	}
}

func (repo LoginMysqlRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	row, err := repo.db.Query("select id, passwd from users where email = ?", email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Passwd); err != nil {
			return nil, err
		}
	}

	return &user, nil
}
