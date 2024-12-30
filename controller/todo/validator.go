package todo

import (
	"fmt"
	"go-to-do/services/todosvc"

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
