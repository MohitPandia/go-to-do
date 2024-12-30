package todosvc

import (
	"fmt"
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/* -------------------------------------------------------------------------- */
/*                                 CreateUser                                 */
/* -------------------------------------------------------------------------- */
func (t *TodoSvcImpl) CreateTodo(ctx *gin.Context, reqBody CreateTodoObject) (models.BaseResponse, entities.Todos, error) {

	fmt.Println("yo coming in service")
	var todo entities.Todos
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	todo.UserPID = reqBody.UserPID
	todo.CategoryPID = reqBody.CategoryPID
	todo.Title = reqBody.Title
	todo.Description = reqBody.Description
	todo.Completed = reqBody.Completed
	todo.DueDate = reqBody.DueDate

	// Save the user in the database
	todo, err = t.todosGorm.CreateTodo(ctx, todo)
	if err != nil {
		return baseRes, todo, errors.Wrap(err, "[CreateTodo][SaveTodo]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "Todo created successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, todo, nil
}
