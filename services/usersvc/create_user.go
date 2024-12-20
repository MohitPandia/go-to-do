package usersvc

import (
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserRes struct {
	BaseRes models.BaseResponse
	User    entities.Users
}

/* -------------------------------------------------------------------------- */
/*                                 CreateUser                                 */
/* -------------------------------------------------------------------------- */
func (u *UserSvcImpl) CreateUser(ctx *gin.Context, reqBody CreateUserObject) (models.BaseResponse, entities.Users, error) {
	var user entities.Users
	var baseRes models.BaseResponse
	var err error

	// Initialize base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	user.Name = reqBody.Name
	user.Email = reqBody.Email
	user.Password = reqBody.Password

	// Save the user in the database
	user, err = u.usersGorm.CreateUser(ctx, user)
	if err != nil {
		return baseRes, user, errors.Wrap(err, "[CreateUser][SaveUser]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "knowledge base created successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, user, nil
}
