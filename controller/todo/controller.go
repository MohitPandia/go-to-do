package todo

import (
	"fmt"
	"go-to-do/smerrors"
	"go-to-do/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                 CreateTodo                                 */
/* -------------------------------------------------------------------------- */
func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	// validate
	fmt.Println("this is me coming in controller before validator")
	reqBody, err := validateTodo(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	fmt.Println("this is me coming in controller after validator")
	// response
	baseRes, user, err := t.todoSvc.CreateTodo(ctx, reqBody)
	if err != nil {
		smerrors.InternalServer(ctx, err.Error())
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	finalRes := createTodoTransformer(baseRes, user)
	utils.ReturnJSONStruct(ctx, finalRes)
}

/* -------------------------------------------------------------------------- */
/*                                 GetAllTodo                                 */
/* -------------------------------------------------------------------------- */

func (t *todoHandler) GetAllTodos(ctx *gin.Context) {
	// Validate the request
	// Validate the request
	page, limit, err := validateGetAllTodos(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	// Call the service method
	baseRes, todos, err := t.todoSvc.GetAllTodos(ctx, page, limit)
	if err != nil {
		smerrors.InternalServer(ctx, err.Error())
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response if needed
	finalRes := GetAllTodoTransformer(baseRes, todos)
	utils.ReturnJSONStruct(ctx, finalRes)
}
