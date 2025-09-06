package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// это как моделька в Laravel
type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
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

func (s *PostRepository) GetById(ctx context.Context, id int64) (*Post, error) {

	query := `
		SELECT id, user_id, title, content, tags, created_at, updated_at 
		FROM posts 
		WHERE id = $1;
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, err
}

func (s *PostRepository) Delete(ctx context.Context, postID int64) error {

	query := "DELETE FROM posts WHERE id = $1"

	res, err := s.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostRepository) Update(ctx context.Context, post *Post) (*Post, error) {
	query := `
        UPDATE posts 
        SET title = $1, content = $2, tags = $3, updated_at = NOW()
        WHERE id = $4
        RETURNING id, user_id, title, content, tags, created_at, updated_at
    `

	var updatedPost Post
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.ID,
	).Scan(
		&updatedPost.ID,
		&updatedPost.UserID,
		&updatedPost.Title,
		&updatedPost.Content,
		pq.Array(&updatedPost.Tags),
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedPost, nil
}
