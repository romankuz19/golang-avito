package repository

import (
	"github.com/jmoiron/sqlx"
	avitoproj "github.com/romankuz19/avito-proj"
)

type Section interface {
	Create(name string) error
	Delete(name string) error
	AddUser(sectionsAdd []string, sectionsDelete []string, userId int) error
	GetUserSections(id int) []string
}

type User interface {
	Create(name string) error
}

type History interface {
	GetUserHistory(userId int, date string) []avitoproj.UserHistory
}

type Repository struct {
	Section
	User
	History
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Section: NewSectionPostgres(db),
		User:    NewUserPostgres(db),
		History: NewHistoryPostgres(db),
	}
}
