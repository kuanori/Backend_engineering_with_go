package repository

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserRepository struct {
	db *sql.DB
}

func (s *UserRepository) Create(ctx context.Context, user *User) error {

	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURN id, created_at,
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
