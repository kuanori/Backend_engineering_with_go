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
	Version   int64     `json:"version"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	User      User      `json"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int `json:"commens_count"`
}

type PostRepository struct {
	db *sql.DB
}

var ErrConflict = errors.New("edit conflict")

func (s *PostRepository) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {

	query := `
		SELECT 
			p.id, p.user_id, p.title, p.content, p.created_at, p.tags, 
			u.username,
			COUNT (c.id) AS comments_count
		FROM posts p
		LEFT JOIN comments c ON c.post_id = p.id
		LEFT JOIN users u ON u.id = p.user_id
		JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
		WHERE f.user_id = $1
		GROUP BY p.id, u.username
		ORDER BY p.created_at ` + fq.Sort + `
		LIMIT $2 OFFSET $3;
		`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feed []PostWithMetadata
	for rows.Next() {
		var p PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			pq.Array(&p.Tags),
			&p.User.Username,
			&p.CommentsCount,
		)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return feed, nil

}

func (s *PostRepository) Create(ctx context.Context, post *Post) error {

	query := `
		INSERT INTO posts (user_id, title, content, tags, version)
		VALUES ($1, $2, $3, $4, 1)
		RETURNING id, created_at, version;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
		&post.Version,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostRepository) GetById(ctx context.Context, id int64) (*Post, error) {

	var post Post
	query := `
	SELECT id, user_id, title, content, tags, version, created_at, updated_at 
	FROM posts 
	WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.Version,
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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
		SET title = $1, content = $2, tags = $3, version = version + 1, updated_at = NOW()
		WHERE id = $4 AND version = $5
		RETURNING id, user_id, title, content, tags, version, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var updatedPost Post
	if err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.ID,
		post.Version, // старое значение
	).Scan(
		&updatedPost.ID,
		&updatedPost.UserID,
		&updatedPost.Title,
		&updatedPost.Content,
		pq.Array(&updatedPost.Tags),
		&updatedPost.Version,
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrEditConflict
		default:
			return nil, err
		}
	}

	return &updatedPost, nil
}
