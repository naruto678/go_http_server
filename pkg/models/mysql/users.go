package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/server-practice/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (user *UserModel) InsertUser(name, email, password string) error {

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `insert into users (name, email, hashed_password, created) values (?,?,?, UTC_TIMESTAMP())`
	_, err = user.DB.Exec(stmt, name, email, hashed_password)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return models.ErrDuplicateMail
			}
		}
	}

	return nil
}

func (user *UserModel) GetUser(id int) (*models.User, error) {
	model := &models.User{}
	stmt := `select id, name , email from users where id = ?`
	err := user.DB.QueryRow(stmt, id).Scan(&model.ID, &model.Name, &model.Email)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return model, err
}

func (user *UserModel) Authenticate(email, password string) (int, error) {
	stmt := `select id, hashed_password from users where email=?`
	row := user.DB.QueryRow(stmt, email)

	var hashed_password []byte
	var id int
	err := row.Scan(&id, &hashed_password)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredential
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredential
	} else if err != nil {
		return 0, err
	}
	return id, nil
}
