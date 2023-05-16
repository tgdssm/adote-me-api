package repositories

import (
	"api/helpers"
	"api/internal/core/domain"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository() *UserMysqlRepository {
	db, err := sql.Open("mysql", helpers.ConnectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return &UserMysqlRepository{
		db: db,
	}
}

func (repo UserMysqlRepository) Create(user *domain.User) (*domain.User, error) {
	statement, err := repo.db.Prepare("insert into users(username, email, cellphone, passwd) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email, user.Cellphone, user.Passwd)
	if err != nil {
		return nil, err
	}

	lastId, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = uint64(lastId)

	return user, nil
}

func (repo UserMysqlRepository) List(queryParameter string) ([]domain.User, error) {
	var users []domain.User
	var rows *sql.Rows
	var err error
	if queryParameter == "" {

		rows, err = repo.db.Query("select u.id, u.username, u.email, u.cellphone, u.created_at, p.id, p.file_name, p.file_path from users u inner join profile_images p on u.id = p.user_id")

	} else {
		queryParameter = fmt.Sprintf("%%%s%%", queryParameter)
		rows, err = repo.db.Query("select u.id, u.username, u.email, u.cellphone, u.created_at, p.id, p.file_name, p.file_path from users u inner join profile_images p on u.id = p.user_id and u.username like ?", queryParameter)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Cellphone, &user.CreatedAt, &user.ProfileImage.ID, &user.ProfileImage.FileName, &user.ProfileImage.FilePath); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo UserMysqlRepository) Get(id int) (*domain.User, error) {
	var user domain.User

	row, err := repo.db.Query("select id, username, email, cellphone, created_at from users where id = ?", id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Cellphone, &user.CreatedAt); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (repo UserMysqlRepository) Update(user *domain.User) (*domain.User, error) {
	statement, err := repo.db.Prepare("update users set username = ?, email = ?, cellphone = ? where id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Cellphone, user.ID)
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
