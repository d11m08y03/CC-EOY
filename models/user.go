package models

import (
	"database/sql"
	"errors"

	"github.com/d11m08y03/CC-EOY/database"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindUserByEmail(email string) (*User, error) {
	row := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)
	user := &User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func CreateUser(user User) error {
	_, err := database.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}