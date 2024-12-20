package user

import (
	"go-to-do/entities"
	"go-to-do/models"
	"go-to-do/services/usersvc"
)

func createUserTransformer(baseRes models.BaseResponse, user entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes usersvc.CreateUserObject

	dataRes.Name = user.Name
	dataRes.Email = user.Email
	dataRes.Name = user.Name
	dataRes.CreatedAt = int(user.CreatedAt.Unix())
	dataRes.UpdatedAt = int(user.UpdatedAt.Unix())

	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes
}
