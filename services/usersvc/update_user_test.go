package usersvc

import (
	"fmt"
	"go-to-do/configs"
	"go-to-do/database"
	"go-to-do/dbops/gorm/users"
	"go-to-do/entities"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	configs.Loadconfigs()

	// Initialize database and services
	gormDB, _ := database.Connection()
	fmt.Println("Database connection established for testing.")

	userGorm := users.Gorm(gormDB)
	userHandler := Handler(userGorm)

	// Create a test user to update
	testUser := entities.Users{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "testpassword",
		PID:      "testpid123",
	}

	// Insert test user into the database
	err := gormDB.Create(&testUser).Error
	assert.Nil(t, err)

	// Set up HTTP context and request body
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := map[string]interface{}{
		"pid": "testpid123",
		"name":     "Updated Test User",
		"email":    "updateduser@example.com",
		"password": "updatedpassword",
	}

	c.Set("Content-Type", "application/json")

	// Call the UpdateUser function
	baseRes, updatedUser, err := userHandler.UpdateUser(c, reqBody)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 200, baseRes.StatusCode)
	assert.Equal(t, "User updated successfully", baseRes.Message)
	assert.Equal(t, "Updated Test User", updatedUser.Name)
	assert.Equal(t, "updateduser@example.com", updatedUser.Email)

	// Cleanup test data
	t.Cleanup(func() {
		gormDB.Model(&entities.Users{}).
			Where("pid = ?", "testpid123").
			Delete(&entities.Users{})
		fmt.Println("Test user deleted successfully.")
	})
}
