package usersvc

import (
	"go-to-do/entities"
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/* -------------------------------------------------------------------------- */
/*                                 DeleteUser                                 */
/* -------------------------------------------------------------------------- */
func (u *UserSvcImpl) DeleteUser(ctx *gin.Context, reqBody DeleteUserObject) (models.BaseResponse, entities.Users, error) {
	var user entities.Users
	var baseRes models.BaseResponse
	var err error

	// Initialize the base response
	baseRes.Success = false
	baseRes.Message = "something went wrong"
	baseRes.StatusCode = http.StatusInternalServerError

	// Call the DB operation to delete the user by PID
	user, err = u.usersGorm.DeleteUser(ctx, reqBody.PID)
	if err != nil {
		return baseRes, user, errors.Wrap(err, "[DeleteUser][DeleteUser]")
	}
	// Set the response
	baseRes.Success = true
	baseRes.Message = "user deleted successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, user, nil
}
