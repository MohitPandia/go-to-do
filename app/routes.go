package app

import (
	"fmt"
	"go-to-do/controller/user"
	"go-to-do/database"
	"go-to-do/dbops/gorm/users"
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

	// orgGorm := orgs.Gorm(gormDB)
	// orgAccountsGorm := orgaccounts.Gorm(gormDB)
	router.POST("/api/users/create", userHandler.CreateUser)
	// router.POST("/api/users/getALl", userHandler.ListUsers)
	// router.POST("/api/users/getUserbyId", userHandler.GetUserByPID)
	// router.POST("/api/users/delete", userHandler.DeleteUser)

	err := router.Run(fmt.Sprintf(":%d", 8080)) // config
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}

}

// func main() {
// 	r := gin.Default()

// 	// PostgreSQL DSN (Data Source Name)
// 	dsn := "host=localhost user=youruser password=yourpassword dbname=tododb sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	// Migrate the tables
// 	if err := db.AutoMigrate(&tables.User{}, &tables.Category{}, &tables.Todo{}); err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	// Routes
// 	r.POST("/api/users", createUser(db))
// 	r.POST("/api/categories", createCategory(db))
// 	r.POST("/api/todos", createTodoItem(db))

// 	r.Run(":8080") // Start the server
// }

// // Handler for creating a user
// func createUser(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user tables.User
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Create(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"status": "User created successfully", "user": user})
// 	}
// }

// // Handler for creating a category
// func createCategory(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var category tables.Category
// 		if err := c.ShouldBindJSON(&category); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Create(&category).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create category"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"status": "Category created successfully", "category": category})
// 	}
// }

// // Handler for creating a todo item
// func createTodoItem(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var todo tables.Todo
// 		if err := c.ShouldBindJSON(&todo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Create(&todo).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create todo item"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"status": "Todo item created successfully", "todo": todo})
// 	}
// }
