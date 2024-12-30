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

	// Extract pagination parameters
	page, limit := reqBody.Page, reqBody.Limit
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Retrieve paginated users
	users, totalCount, err := u.usersGorm.ListUsers(ctx, offset, limit)
	if err != nil {
		return baseRes, nil, errors.Wrap(err, "[ListUsers][ListUsers]")
	}

	// Set the response
	baseRes.Success = true
	baseRes.Message = "users retrieved successfully"
	baseRes.StatusCode = http.StatusOK
	baseRes.Data = users
	baseRes.MetaData = map[string]interface{}{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
	}

	return baseRes, users, nil
}
