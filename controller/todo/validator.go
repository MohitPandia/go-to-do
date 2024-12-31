package todo

import (
	"fmt"
	"go-to-do/services/todosvc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateTodo(c *gin.Context) (todosvc.CreateTodoObject, error) {
	var reqBody todosvc.CreateTodoObject
	var err error

	fmt.Println("comin in validator")
	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	return reqBody, err
}

func validateGetAllTodos(ctx *gin.Context) (int, int, error) {
	// Default values for pagination
	page := 1
	limit := 10

	// Extract and validate "page"
	pageStr := ctx.Query("page")
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil || parsedPage <= 0 {
			return 0, 0, fmt.Errorf("invalid 'page' parameter, must be a positive integer")
		}
		page = parsedPage
	}

	// Extract and validate "limit"
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			return 0, 0, fmt.Errorf("invalid 'limit' parameter, must be a positive integer")
		}
		limit = parsedLimit
	}

	// Additional validation (if required)
	if limit > 100 {
		return 0, 0, fmt.Errorf("'limit' parameter cannot exceed 100")
	}

	return page, limit, nil
}
