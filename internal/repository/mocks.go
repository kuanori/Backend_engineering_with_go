package repository

import (
	"context"
	"database/sql"
	"time"
)

func NewMockRepository() Repository {
	return Repository{
		Users: &MockUserRepository{},
	}
}

type MockUserRepository struct {
}

func (m *MockUserRepository) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockUserRepository) GetById(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}

func (m *MockUserRepository) GetByEmail(context.Context, string) (*User, error) {
	return &User{}, nil
}

func (m *MockUserRepository) CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error {
	return nil
}

func (m *MockUserRepository) Activate(ctx context.Context, t string) error {
	return nil
}

func (m *MockUserRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
