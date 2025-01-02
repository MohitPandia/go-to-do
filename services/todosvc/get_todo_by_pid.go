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
/*                                GetTodoByPID                                */
/* -------------------------------------------------------------------------- */
func (u *TodoSvcImpl) GetTodoByPID(ctx *gin.Context, reqBody GetTodoByPIDObject) (models.BaseResponse, entities.Todos, error) {
	var todo entities.Todos
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	// Fetch the user by PID
	todo, err = u.todosGorm.GetTodoByPID(ctx, reqBody.PID)

	fmt.Println("todo", todo)
	if err != nil {
		return baseRes, todo, errors.Wrap(err, "[GetUserByPID][FetchUser]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "todo fetched successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, todo, nil
}
