// usersvc/service.go
package usersvc

import (
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (u *UserSvcImpl) GetAllUsers(ctx *gin.Context, reqBody GetAllUserObject) (models.BaseResponse, []entities.Users, error) {
	var users []entities.Users
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	// Retrieve all users from the database
	users, err = u.usersGorm.ListUsers(ctx)
	if err != nil {
		return baseRes, nil, errors.Wrap(err, "[ListUsers][ListUsers]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "users retrieved successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, users, nil
}
