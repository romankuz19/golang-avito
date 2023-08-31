package service

import "github.com/romankuz19/avito-proj/pkg/repository"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name string) error {
	return s.repo.Create(name)
}
