package repository

import (
	"context"
	"database/sql"
)

type Follower struct {
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
	CreatedAt  int64 `json:"created_at"`
}

type FollowerRepository struct {
	db *sql.DB
}

func (s *FollowerRepository) Follow(ctx context.Context, followerID, userID int64) error {

	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)
		`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}

func (s *FollowerRepository) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := `
		DELETE FROM followers
		WHERE user_id = $1
		  AND follower_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}
