package cache

import (
	"app/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	rdb *redis.Client
}

const UserExpTime = time.Minute

func (s *UserRepository) Get(ctx context.Context, userID int64) (*repository.User, error) {
	cacheKey := fmt.Sprintf("user-%v", userID)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user repository.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserRepository) Set(ctx context.Context, user *repository.User) error {
	cacheKey := fmt.Sprintf("user-%v", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, cacheKey, json, UserExpTime).Err()
}
