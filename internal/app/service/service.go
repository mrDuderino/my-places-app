package service

import "github.com/mrDuderino/my-places-app/internal/app/repository"

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}
