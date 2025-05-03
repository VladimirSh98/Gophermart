package auth

import (
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
)

type Handler struct {
	User userService.ServiceInterface
}

func NewHandler(user userService.ServiceInterface,
) *Handler {
	return &Handler{User: user}
}
