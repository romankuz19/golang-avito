package repository

import "github.com/jmoiron/sqlx"

type Section interface {
	Create(name string) error
	Delete(name string) error
	AddUser(sectionsAdd []string, sectionsDelete []string, userId int) error
	GetUserSections(id int) []string
}

type User interface {
	Create(name string) error
}

type Repository struct {
	Section
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Section: NewSectionPostgres(db),
		User:    NewUserPostgres(db),
	}
}
