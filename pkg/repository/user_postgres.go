package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(name string) error {
	if name != "" {
		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", usersTable)
		_, err := r.db.Exec(query, name)

		return err
	} else {
		return fmt.Errorf("Empty name")
	}
}
