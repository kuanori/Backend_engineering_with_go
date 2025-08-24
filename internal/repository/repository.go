package repository

import (
	"context"
	"database/sql"
)

// Это больше похоже на объект в PHP с двумя свойствами
type Repository struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewRepository(db *sql.DB) Repository {
	// создает и возвращает экземпляр структуры Repository
	return Repository{
		Posts: &PostRepository{db}, // создается экземпляр структуры PostsRepository
		Users: &UserRepository{db},
	}
}
