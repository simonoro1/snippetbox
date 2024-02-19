package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int
	Name          string
	email         string
	HasedPassword []byte
	Created       time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, pasword string) error {
	return nil
}

func (m *UserModel) Authenticate(email, passsword string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
