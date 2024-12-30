package users

import (
	"fmt"
	"go-to-do/constants"
	"go-to-do/entities"
	"go-to-do/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type GormInterface interface {
	CreateUser(ctx *gin.Context, user entities.Users) (entities.Users, error)
	ListUsers(ctx *gin.Context, offset, limit int) ([]entities.Users, int64, error)
	GetUserByPID(ctx *gin.Context, pid string) (entities.Users, error)
	DeleteUser(ctx *gin.Context, pid string) (entities.Users, error)
	UpdateUser(ctx *gin.Context, pid string, updatedUser entities.Users) (entities.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *userGormImpl {
	return &userGormImpl{DB: gormDB}
}

type userGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                 CreateUser                                 */
/* -------------------------------------------------------------------------- */
func (g *userGormImpl) CreateUser(ctx *gin.Context, user entities.Users) (entities.Users, error) {
	user.PID = utils.UUIDWithPrefix(constants.Prefix.USER)
	err := g.DB.Session(&gorm.Session{}).Create(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "failed to create user")
	}
	return user, nil
}

/* -------------------------------------------------------------------------- */
/*                             List all users.                                */
/* -------------------------------------------------------------------------- */
func (g *userGormImpl) ListUsers(c *gin.Context, offset, limit int) ([]entities.Users, int64, error) {
	var users []entities.Users
	var totalCount int64

	// Count total users
	if err := g.DB.Model(&entities.Users{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to count users")
	}

	// Fetch paginated users
	if err := g.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to list users with pagination")
	}

	return users, totalCount, nil
}

/* -------------------------------------------------------------------------- */
/*                                GetUserByIḌ                                */
/* -------------------------------------------------------------------------- */
func (g *userGormImpl) GetUserByPID(c *gin.Context, pid string) (entities.Users, error) {
	var user entities.Users

	// Query the database
	if err := g.DB.Where("user_pid = ?", pid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by PID")
	}

	fmt.Println("Retrieved user:", user)
	return user, nil
}

/* -------------------------------------------------------------------------- */
/*                                 DeleteUser                                 */
/* -------------------------------------------------------------------------- */
func (g *userGormImpl) DeleteUser(c *gin.Context, pid string) (entities.Users, error) {
	var user entities.Users
	fmt.Println("In gorm method: PID", pid)

	// Check if the user exists
	if err := g.DB.Where("user_pid = ?", pid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by PID")
	}

	// Proceed with deletion if user exists
	if err := g.DB.Where("user_pid = ?", pid).Delete(&user).Error; err != nil {
		fmt.Println("Error deleting user:", err)
		return user, errors.Wrap(err, "failed to delete user")
	}

	fmt.Println("User deleted successfully:", user)
	return user, nil
}

/* -------------------------------------------------------------------------- */
/*                                 UpdateUser                                 */
/* -------------------------------------------------------------------------- */
func (g *userGormImpl) UpdateUser(ctx *gin.Context, pid string, updatedUser entities.Users) (entities.Users, error) {
	var user entities.Users

	// Check if the user exists
	if err := g.DB.Where("user_pid = ?", pid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by PID")
	}

	// Update the fields using the provided user object
	if err := g.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		return user, errors.Wrap(err, "failed to update user")
	}

	fmt.Println("User updated successfully:", user)
	return user, nil
}
