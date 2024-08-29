package handlers

import "github.com/leal/pkg/users/domain"

type Handler struct {
	UserService domain.UserService
}

func NewUserHandler(s domain.UserService) *Handler {
	return &Handler{
		UserService: s,
	}
}
