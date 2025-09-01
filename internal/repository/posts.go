package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// это как моделька в Laravel
type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostRepository struct {
	db *sql.DB
}

func (s *PostRepository) Create(ctx context.Context, post *Post) error {

	query := `
		INSERT INTO posts (user_id, title, content, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at;
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
