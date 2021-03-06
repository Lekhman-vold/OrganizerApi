package repository

import (
	org "github.com/Lekhman-vold/OrganizerApi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user org.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
