package app

import (
	"fmt"
	"go-to-do/controller/todo"
	"go-to-do/controller/user"
	"go-to-do/database"
	"go-to-do/dbops/gorm/todos"
	"go-to-do/dbops/gorm/users"
	"go-to-do/services/todosvc"
	"go-to-do/services/usersvc"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
)

// skipping logs for "/" routes
func Default() *gin.Engine {
	engine := gin.New()
	conf := gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}
	engine.Use(gin.LoggerWithConfig(conf), gin.Recovery())
	return engine
}

func MapUrl() {
	fmt.Println("so also came in routes")
	router := Default()

	gormDB, _ := database.Connection()
	fmt.Println("connection", gormDB)

	userGorm := users.Gorm(gormDB)
	userSvc := usersvc.Handler(userGorm)
	userHandler := user.Handler(userSvc)

	router.POST("/api/users/create", userHandler.CreateUser)
	router.GET("/api/users/getAll", userHandler.GetAllUsers)
	router.GET("/api/users/get-by-pid", userHandler.GetUserByPID)
	router.POST("/api/users/delete", userHandler.DeleteUser)
	router.PUT("api/users/update", userHandler.UpdateUser)

	todoGorm := todos.Gorm(gormDB)
	todosvc := todosvc.Handler(todoGorm)
	todoHandler := todo.Handler(todosvc)

	router.POST("/api/todo/create", todoHandler.CreateTodo)
	// router.PUT("api/todo/update", todoHandler.UpdateTodo)
	router.GET("api/todo/getAllTodo", todoHandler.GetAllTodos)
	// router.GET("api/todo/getTodo-by-pid", todoHandler.GetTodoByPid)
	// router.POST("/api/todo/delete", todoHandler.deleteTodo)

	err := router.Run(fmt.Sprintf(":%d", 8080)) // config
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}

}
