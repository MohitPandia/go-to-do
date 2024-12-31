package todosvc

import (
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (t *TodoSvcImpl) GetAllTodos(ctx *gin.Context, page int, limit int) (models.BaseResponse, []entities.Todos, error) {
	var todos []entities.Todos
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Retrieve paginated todos
	todos, totalCount, err := t.todosGorm.GetAllTodos(ctx, offset, limit)
	if err != nil {
		return baseRes, nil, errors.Wrap(err, "[ListTodo][ListTodo]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "users retrieved successfully"
	baseRes.StatusCode = http.StatusOK
	baseRes.Data = todos
	baseRes.MetaData = map[string]interface{}{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
	}

	return baseRes, todos, nil
}
