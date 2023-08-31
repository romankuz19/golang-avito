package service

import (
	"github.com/romankuz19/avito-proj/pkg/repository"
)

type SectionService struct {
	repo repository.Section
}

func NewSectionService(repo repository.Section) *SectionService {
	return &SectionService{repo: repo}
}

func (s *SectionService) Create(name string) error {
	return s.repo.Create(name)
}

func (s *SectionService) Delete(name string) error {
	return s.repo.Delete(name)
}

func (s *SectionService) AddUser(sectionsAdd []string, sectionsDelete []string, userId int) error {
	return s.repo.AddUser(sectionsAdd, sectionsDelete, userId)
}

func (s *SectionService) GetUserSections(id int) []string {
	return s.repo.GetUserSections(id)
}
