package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type userMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository() *userMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return &userMysqlRepository{
		db: db,
	}
}

func (repo userMysqlRepository) Create(user *domain.User) (*domain.User, error) {
	statement, err := repo.db.Prepare("insert into users(username, email, cellphone, passwd, picturePath) values(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Cellphone, user.Passwd, user.PicturePath)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = uint64(lastId)

	return user, nil
}

func (repo userMysqlRepository) List(queryParameter string) ([]domain.User, error) {
	users := []domain.User{}
	var rows *sql.Rows
	var err error
	if queryParameter == "" {
		rows, err = repo.db.Query("select id, username, email, picturePath, cellphone, createdAt from users")

	} else {
		queryParameter = fmt.Sprintf("%%%s%%", queryParameter)
		rows, err = repo.db.Query("select id, username, email, picturePath, cellphone, createdAt from users where username like ?", queryParameter)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.PicturePath, &user.Cellphone, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo userMysqlRepository) Get(id int) (*domain.User, error) {
	var user domain.User

	row, err := repo.db.Query("select id, username, email, picturePath, cellphone, createdAt from users where id = ?", id)

	if err != nil {
		return nil, err
	}

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.PicturePath, &user.Cellphone, &user.CreatedAt); err != nil {
			return nil, err
		}
	}

	return &user, nil
}
