package services

import "github.com/leal/pkg/users/domain"

var _ domain.UserService = &Service{}

type Service struct {
	Repository domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *Service {
	return &Service{
		Repository: repo,
	}
}
