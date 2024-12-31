package todosvc

import (
	"go-to-do/dbops/gorm/todos"
	"go-to-do/entities"
	"go-to-do/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  Reciever                                  */
/* -------------------------------------------------------------------------- */
type TodoSvcImpl struct {
	todosGorm todos.GormInterface
}

/* -------------------------------------------------------------------------- */
/*                              Service Interface                             */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreateTodo(ctx *gin.Context, reqBody CreateTodoObject) (models.BaseResponse, entities.Todos, error)
	GetAllTodos(ctx *gin.Context, page int, limit int) (models.BaseResponse, []entities.Todos, error)
	GetTodoByPID(ctx *gin.Context, reqBody GetTodoByPIDObject) (models.BaseResponse, entities.Todos, error)
}

/* -------------------------------------------------------------------------- */
/*                               HANDLER                                      */
/* -------------------------------------------------------------------------- */
func Handler(todosGorm todos.GormInterface) *TodoSvcImpl {
	return &TodoSvcImpl{
		todosGorm: todosGorm,
	}
}
