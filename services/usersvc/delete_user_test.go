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

func TestDeleteUser(t *testing.T) {
	configs.Loadconfigs()

	// Initialize database and services
	gormDB, _ := database.Connection()
	fmt.Println("Database connection established for testing.")

	userGorm := users.Gorm(gormDB)
	userHandler := Handler(userGorm)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Add a test user
	testUser := entities.Users{
		PID:   "test-user-003",
		Name:  "Test User",
		Email: "testuser003@example.com",
	}
	err := gormDB.Create(&testUser).Error
	assert.Empty(t, err)

	// Prepare request for DeleteUser
	reqBody := DeleteUserObject{
		PID: testUser.PID,
	}

	// Call DeleteUser handler
	c.Request = &http.Request{
		Header: make(http.Header),
		Method: http.MethodDelete,
	}
	c.Params = append(c.Params, gin.Param{Key: "pid", Value: reqBody.PID})

	baseRes, _, err := userHandler.DeleteUser(c, reqBody)
	assert.Empty(t, err)
	assert.NotEmpty(t, baseRes)
	assert.Equal(t, true, baseRes.Success)
	assert.Equal(t, 200, baseRes.StatusCode)

	// Verify the user has been deleted
	var deletedUser entities.Users
	result := gormDB.Where("user_pid = ?", reqBody.PID).First(&deletedUser)
	assert.NotEmpty(t, result.Error)
	assert.True(t, gormDB.RecordNotFound(result.Error))

	// Cleanup test user data
	t.Cleanup(func() {
		// Ensure the test user is deleted
		gormDB.Model(&entities.Users{}).
			Where("user_pid = ?", reqBody.PID).
			Delete(&entities.Users{})
		fmt.Println("Test user deleted successfully.")
	})
}
