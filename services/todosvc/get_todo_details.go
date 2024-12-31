package todosvc

// import (
// 	"go-to-do/entities"
// 	"go-to-do/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pkg/errors"
// )

// func (t *TodoSvcImpl) GetTodoDetails(ctx *gin.Context) (models.BaseResponse, entities.Todos, entities.Users, error) {
// 	var todo entities.Todos
// 	var user entities.Users
// 	var err error
// 	var baseRes models.BaseResponse

// 	// Initialize base response
// 	baseRes.Success = false
// 	baseRes.Message = "something went wrong"
// 	baseRes.StatusCode = http.StatusInternalServerError

// 	todo, err = t.todosGorm.GetAllTodos(ctx)
// 	if err != nill {
// 		return baseRes, todo, user, errors.Wrap(err, "[FindUser]")
// 	}

// }
