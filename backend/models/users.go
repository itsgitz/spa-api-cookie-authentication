package models

import (
	"context"
	"database/sql"
	"go-api/database"
)

type Users struct {
	ID       uint64            `json:"id,omitempty"`
	Username string            `json:"username,omitempty"`
	Password string            `json:"password,omitempty"`
	Database database.Database `json:"-"`
}

func (u *Users) FindUsers(ctx context.Context) (*[]Users, error) {
	db, err := u.Database.Connection()
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

	// NOTE: this is just an example
	// The password is not encrypted, its only plain string
	for rows.Next() {
		var u Users
		if err := rows.Scan(&u.ID, &u.Username, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &users, nil
}

func (u *Users) Authenticate(ctx context.Context) (bool, error) {
	db, err := u.Database.Connection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var usernameFromTable string = ""

	// NOTE: this is just an example
	// The password is not encrypted, its only plain string
	row := db.QueryRowContext(
		ctx,
		"SELECT username FROM users WHERE username = $1 AND password = $2",
		u.Username,
		u.Password,
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
