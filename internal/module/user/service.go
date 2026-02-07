package user

import "github.com/Alwin18/golang-modular-template/internal/shared/logger"

type Service interface {
	GetUsers() ([]User, error)
}

type service struct {
	repo   Repository
	logger logger.Logger
}

func NewService(r Repository, l logger.Logger) Service {
	return &service{r, l}
}

func (s *service) GetUsers() ([]User, error) {
	s.logger.Info("fetch users")
	return s.repo.FindAll()
}
