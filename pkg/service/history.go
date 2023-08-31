package service

import (
	avitoproj "github.com/romankuz19/avito-proj"
	"github.com/romankuz19/avito-proj/pkg/repository"
)

type HistoryService struct {
	repo repository.History
}

func NewHistoryService(repo repository.History) *HistoryService {
	return &HistoryService{repo: repo}
}

func (s *HistoryService) GetUserHistory(userId int, date string) []avitoproj.UserHistory {
	return s.repo.GetUserHistory(userId, date)
}
