package usersvc

import (
	"fmt"
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (u *UserSvcImpl) GetUserByPID(ctx *gin.Context, reqBody GetUserByPIDObject) (models.BaseResponse, entities.Users, error) {
	var user entities.Users
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	// Fetch the user by PID
	user, err = u.usersGorm.GetUserByPID(ctx, reqBody.PID)

	fmt.Println("user", user)
	if err != nil {
		return baseRes, user, errors.Wrap(err, "[GetUserByPID][FetchUser]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "user fetched successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, user, nil
}
