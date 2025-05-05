package auth

import (
	operationService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operation"
)

type Handler struct {
	Operation operationService.ServiceInterface
}

func NewHandler(operation operationService.ServiceInterface,
) *Handler {
	return &Handler{Operation: operation}
}
