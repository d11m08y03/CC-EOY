package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/d11m08y03/CC-EOY/database"
	"github.com/d11m08y03/CC-EOY/logger"
)

type Organisor struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  uint   `json:"is_admin"`
}

func FindUserByEmail(email string) (*Organisor, error) {
	row := database.DB.QueryRow("SELECT id, name, email, password FROM organisors WHERE email = ?", email)
	organisor := &Organisor{}
	if err := row.Scan(&organisor.ID, &organisor.Name, &organisor.Email, &organisor.Password); err != nil {
		logger.Error(fmt.Sprintf("Failed to find user %s in DB: %s", organisor.Name, err.Error()))

		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return organisor, nil
}

func CreateOrganisor(user Organisor) error {
	_, err := database.DB.Exec("INSERT INTO organisors (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)

	return err
}
