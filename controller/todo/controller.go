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
