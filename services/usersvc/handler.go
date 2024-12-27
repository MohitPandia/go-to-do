package usersvc

import (
	"go-to-do/dbops/gorm/users"
	"go-to-do/entities"
	"go-to-do/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  Reciever                                  */
/* -------------------------------------------------------------------------- */
type UserSvcImpl struct {
	usersGorm users.GormInterface
}

/* -------------------------------------------------------------------------- */
/*                              Service Interface                             */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreateUser(ctx *gin.Context, reqBody CreateUserObject) (models.BaseResponse, entities.Users, error)
	GetAllUsers(ctx *gin.Context, reqBody GetAllUserObject) (models.BaseResponse, []entities.Users, error)
	GetUserByPID(ctx *gin.Context, reqBody GetUserByPIDObject) (models.BaseResponse, entities.Users, error)
	DeleteUser(ctx *gin.Context, reqBody DeleteUserObject) (models.BaseResponse, entities.Users, error)
	UpdateUser(ctx *gin.Context, reqBody map[string]interface{}) (models.BaseResponse, entities.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                               HANDLER                                      */
/* -------------------------------------------------------------------------- */
func Handler(usersGorm users.GormInterface) *UserSvcImpl {
	return &UserSvcImpl{
		usersGorm: usersGorm,
	}
}
