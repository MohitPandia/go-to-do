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

func GetAllUsertransformer(baseRes models.BaseResponse, users []entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var usersData []entities.Users

	// Map entities.Users to usersvc.GetAllUserObject
	for _, user := range users {
		userData := entities.Users{
			Name:      user.Name,
			Email:     user.Email,
			PID:       user.PID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersData = append(usersData, userData)
	}

	// Populate final response
	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = usersData
	finalRes.MetaData = baseRes.MetaData

	return finalRes
}

func TransformGetUserResponse(baseRes models.BaseResponse, user entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes entities.Users

	dataRes.PID = user.PID
	dataRes.Name = user.Name
	dataRes.Email = user.Email

	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes
}

func DeleteUserTransformer(baseRes models.BaseResponse, user entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes entities.Users

	dataRes.PID = user.PID
	dataRes.Name = user.Name
	dataRes.Email = user.Email

	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes

}

func UpdateUserTransformer(baseRes models.BaseResponse, updatedUser entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes entities.Users

	dataRes.PID = updatedUser.PID
	dataRes.Name = updatedUser.Name
	dataRes.Email = updatedUser.Email
	// Optionally add more fields if you want to include other fields like "Password"

	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes
}
