package usersvc

import (
	// "fmt"
	"go-to-do/entities"
	"go-to-do/models"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (u *UserSvcImpl) UpdateUser(ctx *gin.Context, reqBody map[string]interface{}) (models.BaseResponse, entities.Users, error) {
	var baseRes models.BaseResponse
	var user entities.Users

	// Validate required fields
	pid, ok := reqBody["pid"].(string)
	if !ok || pid == "" {
		baseRes.StatusCode = 400
		return baseRes, entities.Users{}, errors.New("PID is required")
	}

	// Fetch the existing user by PID
	user, err := u.usersGorm.GetUserByPID(ctx, pid)
	if err != nil {
		baseRes.StatusCode = 404
		return baseRes, entities.Users{}, errors.New("user not found")
	}

	// Update the user fields dynamically from reqBody
	if name, ok := reqBody["name"].(string); ok {
		user.Name = name
	}
	if email, ok := reqBody["email"].(string); ok {
		user.Email = email
	}
	if password, ok := reqBody["password"].(string); ok {
		user.Password = password
	}

	// Save the updated user details
	updatedUser, err := u.usersGorm.UpdateUser(ctx, pid, user)
	if err != nil {
		baseRes.StatusCode = 500
		return baseRes, entities.Users{}, errors.New("failed to update user")
	}

	// Populate the updated details in the response
	baseRes.StatusCode = 200
	baseRes.Message = "User updated successfully"
	baseRes.Data = updatedUser
	return baseRes, updatedUser, nil
}
