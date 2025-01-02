package todo

import (
	"go-to-do/entities"
	"go-to-do/models"
	"go-to-do/services/todosvc"
)


/* -------------------------------------------------------------------------- */
/*                            createTodoTransformer                           */
/* -------------------------------------------------------------------------- */
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


/* -------------------------------------------------------------------------- */
/*                            GetAllTodoTransformer                           */
/* -------------------------------------------------------------------------- */
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

/* -------------------------------------------------------------------------- */
/*                       TransformGetTodoDetailsResponse                      */
/* -------------------------------------------------------------------------- */
func TransformGetTodoDetailsResponse(baseRes models.BaseResponse, todo entities.Todos, user entities.Users) models.BaseResponse {
	var finalRes models.BaseResponse
	var dataRes struct {
		TodoID      int    `json:"todo_id"`
		UserName    string `json:"user_name"`
		UserEmail   string `json:"user_email"`
		UserPID     string `json:"user_pid"`
		CategoryPID string `json:"category_pid"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
		DueDate     string `json:"due_date"`
	}

	// Map todo details
	dataRes.TodoID = todo.ID
	dataRes.UserPID = todo.UserPID
	dataRes.CategoryPID = todo.CategoryPID
	dataRes.Title = todo.Title
	dataRes.Description = todo.Description
	dataRes.Completed = todo.Completed
	dataRes.DueDate = todo.DueDate.Format("2006-01-02") // Format date as YYYY-MM-DD

	// Map user details
	dataRes.UserName = user.Name
	dataRes.UserEmail = user.Email

	// Set response metadata
	finalRes.Success = baseRes.Success
	finalRes.Message = baseRes.Message
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Data = dataRes

	return finalRes
}
