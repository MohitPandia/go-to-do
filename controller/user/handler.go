package user

import (
	"go-to-do/services/usersvc"
)

type userHandler struct {
	userSvc usersvc.Interface
}

func Handler(userSvc usersvc.Interface) *userHandler {
	return &userHandler{
		userSvc: userSvc,
	}
}