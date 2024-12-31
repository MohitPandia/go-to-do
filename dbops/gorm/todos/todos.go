package todos

import (
	"fmt"
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
	GetAllTodos(ctx *gin.Context, offset, limit int) ([]entities.Todos, int64, error)
	GetTodoByPID(ctx *gin.Context, pid string) (entities.Todos, error)
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

/* -------------------------------------------------------------------------- */
/*                                 GetAllTodo                                 */
/* -------------------------------------------------------------------------- */
func (g *todosGormImpl) GetAllTodos(c *gin.Context, offset, limit int) ([]entities.Todos, int64, error) {
	var todos []entities.Todos
	var totalCount int64

	// Count total todo
	if err := g.DB.Model(&entities.Todos{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to count todo")
	}

	// Fetch paginated todo
	if err := g.DB.Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to list todo with pagination")
	}
	fmt.Println("totalCount", totalCount)

	return todos, totalCount, nil
}

/* -------------------------------------------------------------------------- */
/*                                GetTodoByPid                                */
/* -------------------------------------------------------------------------- */
func (g *todosGormImpl) GetTodoByPID(c *gin.Context, pid string) (entities.Todos, error) {
	var todo entities.Todos

	// Query the database
	if err := g.DB.Where("todos_pid = ?", pid).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, errors.Wrap(err, "todo not found")
		}
		return todo, errors.Wrap(err, "failed to get user by PID")
	}

	fmt.Println("Retrieved user:", todo)
	return todo, nil
}
