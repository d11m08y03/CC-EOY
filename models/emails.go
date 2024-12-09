package models

import (
	"fmt"

	"github.com/d11m08y03/CC-EOY/database"
)

type Email struct {
	ID          int
	Email       string
	Password    string
	AppPassword string
	Sent        int
}

func GetAllEmails() ([]Email, error) {
	query := "SELECT ID, Email, Password, AppPassword, Sent FROM emails"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var emails []Email
	for rows.Next() {
		var email Email
		if err := rows.Scan(&email.ID, &email.Email, &email.Password, &email.AppPassword, &email.Sent); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return emails, nil
}
