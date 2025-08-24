package repository

import (
	"context"
	"database/sql"
)

// Это больше похоже на объект в PHP с двумя свойствами
type Repository struct {
	Posts interface {
		Create(context.Context) error
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewRepository(db *sql.DB) Repository {
	// создает и возвращает экземпляр структуры Repository
	return Repository{
		Posts: &PostsRepository{db}, // создается экземпляр структуры PostsRepository
		Users: &UsersRepository{db},
	}
}
