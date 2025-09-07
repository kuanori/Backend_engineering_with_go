package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("Resource not found")
	ErrEditConflict      = errors.New("edit conflict")
	QueryTimeoutDuration = time.Second * 5
)

// Это больше похоже на объект в PHP с двумя свойствами
type Repository struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
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
