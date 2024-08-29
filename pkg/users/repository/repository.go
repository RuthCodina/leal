package repository

import (
	"database/sql"

	"github.com/leal/pkg/users/domain"
)

var _ domain.UserRepository = &Repository{}

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}
