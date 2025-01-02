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

/* -------------------------------------------------------------------------- */
/*                                GetTodoByPID                                */
/* -------------------------------------------------------------------------- */
func (t *todoHandler) GetTodoByPID(ctx *gin.Context) {
	// Validate request body
	reqBody, err := validateGetTodoByPID(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	fmt.Println("reqBody", reqBody)
	// Call the service method
	baseRes, todo, err := t.todoSvc.GetTodoByPID(ctx, reqBody)
	if err != nil {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	fmt.Println("baseRes", baseRes)
	fmt.Println("todo", todo)
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response
	finalRes := TransformGetTodoResponse(baseRes, todo)
	utils.ReturnJSONStruct(ctx, finalRes)
}

/* -------------------------------------------------------------------------- */
/*                               GetTodoDetails                               */
/* -------------------------------------------------------------------------- */
func (t *todoHandler) GetTodoDetails(ctx *gin.Context) {

	// Validate request body
	reqBody, err := validateGetTodoByPID(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	fmt.Println("reqBody", reqBody)
	// Call the service method
	baseRes, todo, user, err := t.todoSvc.GetTodoDetails(ctx, reqBody)
	fmt.Println("todo", todo)
	fmt.Println("user", user)
	fmt.Println("baseRes", baseRes)
	if err != nil {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response
	finalRes := TransformGetTodoDetailsResponse(baseRes, todo, user)
	fmt.Println("finalRes", finalRes)
	utils.ReturnJSONStruct(ctx, finalRes)
}
