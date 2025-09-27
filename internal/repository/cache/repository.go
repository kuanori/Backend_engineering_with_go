package cache

import (
	"app/internal/repository"
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Users interface {
		Get(context.Context, int64) (*repository.User, error)
		Set(context.Context, *repository.User) error
	}
}

func NewRedisRepository(rdb *redis.Client) Repository {
	return Repository{
		Users: &UserRepository{rdb: rdb},
	}
}
