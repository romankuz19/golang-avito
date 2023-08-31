package service

import (
	avitoproj "github.com/romankuz19/avito-proj"
	"github.com/romankuz19/avito-proj/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

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

type Service struct {
	Section
	User
	History
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Section: NewSectionService(repos.Section),
		User:    NewUserService(repos.User),
		History: NewHistoryService(repos.History),
	}
}
