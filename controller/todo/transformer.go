package todo

import (
	"go-to-do/entities"
	"go-to-do/models"
	"go-to-do/services/todosvc"
)

func createTodoTransformer(baseRes models.BaseResponse, todo entities.Todos) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes todosvc.CreateTodoObject

	dataRes.UserPID = todo.UserPID
	dataRes.CategoryPID = todo.CategoryPID
	dataRes.Title = todo.Title
	dataRes.Description = todo.Description
	dataRes.Completed = todo.Completed
	dataRes.DueDate = todo.DueDate

	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes
}
