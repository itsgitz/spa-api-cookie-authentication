package models

import (
	"context"
	"go-api/database"
	"time"
)

type Notes struct {
	ID        uint64            `json:"id,omitempty"`
	Content   string            `json:"content,omitempty"`
	CreatedAt time.Time         `json:"created_at,omitempty"`
	Database  database.Database `json:"-"`
}

func (n *Notes) FindNotes(ctx context.Context) (*[]Notes, error) {
	db, err := n.Database.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Notes

	for rows.Next() {
		var n Notes
		if err := rows.Scan(&n.ID, &n.Content, &n.CreatedAt); err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	return &notes, nil
}

func (n *Notes) SaveNotes(ctx context.Context) (int64, error) {
	db, err := n.Database.Connection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Millisecond)
	defer cancel()

	query := `INSERT INTO notes (content) VALUES ($1)`

	result, err := db.ExecContext(ctx, query, n.Content)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
