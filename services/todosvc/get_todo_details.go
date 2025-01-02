package todosvc

import (
	"fmt"
	"go-to-do/database"
	"go-to-do/dbops/gorm/users"
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (t *TodoSvcImpl) GetTodoDetails(ctx *gin.Context, reqBody GetTodoByPIDObject) (models.BaseResponse, entities.Todos, entities.Users, error) {
	var todo entities.Todos
	var user entities.Users
	var err error
	var baseRes models.BaseResponse

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	todo, err = t.todosGorm.GetTodoByPID(ctx, reqBody.PID)
	if err != nil {
		return baseRes, todo, user, errors.Wrap(err, "[GetTodoByPID]")
	}

	var userPid = todo.UserPID

	gormDB, _ := database.Connection()
	usersGorm := users.Gorm(gormDB)
	user, err = usersGorm.GetUserByPID(ctx, userPid)
	if err != nil {
		return baseRes, todo, user, errors.Wrap(err, "[GetUserByPID]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "todo fetched successfully"
	baseRes.StatusCode = http.StatusOK

	fmt.Println("baseRes------------>", baseRes)
	return baseRes, todo, user, nil
}
