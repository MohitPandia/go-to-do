package usersvc

import (
	"fmt"
	"go-to-do/configs"
	"go-to-do/database"
	"go-to-do/dbops/gorm/users"
	"go-to-do/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	configs.Loadconfigs()

	// Initialize database and services
	gormDB, _ := database.Connection()
	fmt.Println("Database connection established for testing.")

	userGorm := users.Gorm(gormDB)
	userHandler := Handler(userGorm)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	var req CreateUserObject

	baseRes, res, err := userHandler.CreateUser(c, req)
	assert.Empty(t, err)
	assert.NotEmpty(t, baseRes)
	assert.NotEmpty(t, res)
	assert.Equal(t, true, baseRes.Success)
	assert.Equal(t, 200, baseRes.StatusCode)

	// Cleanup test data
	t.Cleanup(func() {
		gormDB.Model(&entities.Users{}).
			Where("id = ?", res.PID).
			Delete(&entities.Users{})
		fmt.Println("Test user deleted successfully.")
	})
}
