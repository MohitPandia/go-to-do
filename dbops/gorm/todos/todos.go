package todos

import (
	"go-to-do/constants"
	"go-to-do/entities"
	"go-to-do/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type GormInterface interface {
	CreateTodo(ctx *gin.Context, todos entities.Todos) (entities.Todos, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *todosGormImpl {
	return &todosGormImpl{DB: gormDB}
}

type todosGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                 CreateTodo                                 */
/* -------------------------------------------------------------------------- */
func (g *todosGormImpl) CreateTodo(ctx *gin.Context, todos entities.Todos) (entities.Todos, error) {
	todos.PID = utils.UUIDWithPrefix(constants.Prefix.TODO)
	err := g.DB.Session(&gorm.Session{}).Create(&todos).Error
	if err != nil {
		return todos, errors.Wrap(err, "failed to create todo")
	}
	return todos, nil
}
