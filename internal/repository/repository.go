package repository

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Resource not found")
)

// Это больше похоже на объект в PHP с двумя свойствами
type Repository struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewRepository(db *sql.DB) Repository {
	// создает и возвращает экземпляр структуры Repository
	return Repository{
		Posts:    &PostRepository{db}, // создается экземпляр структуры PostsRepository
		Users:    &UserRepository{db},
		Comments: &CommentRepository{db},
	}
}
