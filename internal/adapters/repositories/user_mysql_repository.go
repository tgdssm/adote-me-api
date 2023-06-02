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
	users := []domain.User{}
	var rows *sql.Rows
	var err error
	if queryParameter == "" {
		rows, err = repo.db.Query("select u.id, u.username, u.email, u.cellphone, u.created_at from users u ")

	} else {
		queryParameter = fmt.Sprintf("%%%s%%", queryParameter)
		rows, err = repo.db.Query("select u.id, u.username, u.email, u.cellphone, u.created_at from users u where u.username like ?", queryParameter)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		//var ID *uint64
		//var fileName *string
		//var filePath *string
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Cellphone, &user.CreatedAt); err != nil {
			return nil, err
		}
		//if ID != nil {
		//	user.ProfileImage.ID = *ID
		//	user.ProfileImage.FileName = *fileName
		//	user.ProfileImage.FilePath = *filePath
		//}

		users = append(users, user)
	}

	return users, nil
}

func (repo UserMysqlRepository) Get(id int) (*domain.User, error) {
	var user domain.User

	row, err := repo.db.Query("select u.id, u.username, u.email, u.cellphone, u.created_at from users u where u.id = ?", id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		//var ID *uint64
		//var fileName *string
		//var filePath *string
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Cellphone, &user.CreatedAt); err != nil {
			return nil, err
		}
		//if ID != nil {
		//	user.ProfileImage.ID = *ID
		//	user.ProfileImage.FileName = *fileName
		//	user.ProfileImage.FilePath = *filePath
		//}
	}

	return &user, nil
}

func (repo UserMysqlRepository) Update(user *domain.User) error {
	statement, err := repo.db.Prepare("update users set username = ?, email = ?, cellphone = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Cellphone, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserMysqlRepository) Delete(id int) error {
	statement, err := repo.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
