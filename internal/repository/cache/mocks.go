package cache

import (
	"app/internal/repository"
	"context"
)

func NewMockRepository() Repository {
	return Repository{
		Users: &MockUserRepository{},
	}
}

type MockUserRepository struct {
}

func (m MockUserRepository) Get(ctx context.Context, id int64) (*repository.User, error) {
	return nil, nil
}

func (m MockUserRepository) Set(ctx context.Context, user *repository.User) error {
	return nil
}
