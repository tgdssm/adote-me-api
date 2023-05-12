package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{
		db,
	}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users(username, email, passwd, picturePath) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Passwd, user.PicturePath)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (u users) GetUsers(queryParameter string) ([]models.User, error) {
	users := []models.User{}
	var rows *sql.Rows
	var err error
	if queryParameter == "" {
		rows, err = u.db.Query("select id, username, email, picturePath, createdAt from users")

	} else {
		queryParameter = fmt.Sprintf("%%%s%%", queryParameter)
		rows, err = u.db.Query("select id, username, email, picturePath, createdAt from users where username like ?", queryParameter)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.PicturePath, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) GetUser(id int) (models.User, error) {
	var user models.User

	row, err := u.db.Query("select id, username, email, picturePath, createdAt from users where id = ?", id)

	if err != nil {
		return models.User{}, err
	}

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.PicturePath, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
