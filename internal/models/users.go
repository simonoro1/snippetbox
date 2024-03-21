package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int
	Name          string
	Email         string
	HasedPassword []byte
	Created       time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, pasword string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(pasword), 12)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
			VALUES(?,?,?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hash))

	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, passsword string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
