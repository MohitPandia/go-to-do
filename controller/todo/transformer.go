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

func GetAllTodoTransformer(baseRes models.BaseResponse, todos []entities.Todos) models.BaseResponse {
	var finalRes models.BaseResponse
	var todosData []entities.Todos

	// Map entities.Todos to todosvc.GetAllTodosObject
	for _, todo := range todos {
		todoData := entities.Todos{
			UserPID:     todo.UserPID,
			PID:         todo.PID,
			CategoryPID: todo.CategoryPID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		}
		todosData = append(todosData, todoData)
	}

	// Populate final response
	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = todosData
	finalRes.MetaData = baseRes.MetaData

	return finalRes
}

/* -------------------------------------------------------------------------- */
/*                          TransformGetTodoResponse                          */
/* -------------------------------------------------------------------------- */

func TransformGetTodoResponse(baseRes models.BaseResponse, todo entities.Todos) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes entities.Todos

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
