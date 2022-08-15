package models

import (
	"context"
	"database/sql"
	"go-api/database"
)

type Users struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *Users) FindUsers(ctx context.Context) (*[]Users, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Users

	for rows.Next() {
		var u Users
		if err := rows.Scan(&u.ID, &u.Username, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &users, nil
}

func (u *Users) Authenticate(ctx context.Context, username, password string) (bool, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var usernameFromTable string = ""

	row := db.QueryRowContext(
		ctx,
		"SELECT username FROM users WHERE username = $1 AND password = $2",
		username,
		password,
	)

	err = row.Scan(&usernameFromTable)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	}

	if usernameFromTable == "" {
		return false, nil
	}

	return true, nil
}
