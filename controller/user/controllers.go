package user

import (
	"go-to-do/smerrors"
	"go-to-do/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. Validate request
// 2. Send Request to service layer
// 3. Get response from service layer
// 4. Send Response back with transformers

/* -------------------------------------------------------------------------- */
/*                                 CreateUser                                 */
/* -------------------------------------------------------------------------- */
func (u *userHandler) CreateUser(ctx *gin.Context) {
	// validate
	reqBody, err := validateCreateUser(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	// response
	baseRes, user, err := u.userSvc.CreateUser(ctx, reqBody)
	if err != nil {
		smerrors.InternalServer(ctx, err.Error())
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	finalRes := createUserTransformer(baseRes, user)
	utils.ReturnJSONStruct(ctx, finalRes)
}



///--->

// func (u *userHandler) ListUsers(ctx *gin.Context) {
// 	// Fetch users from the service
// 	baseRes, users, err := u.userSvc.ListUsers(ctx)
// 	if err != nil {
// 		smerrors.InternalServer(ctx, err.Error())
// 		return
// 	}

// 	utils.ReturnJSONStruct(ctx, baseRes)
// }


// func (u *userHandler) GetUserByPID(ctx *gin.Context) {
// 	// Validate the request body
// 	reqBody, err := validateGetUserById(ctx)
// 	if err != nil {
// 		smerrors.Validation(ctx, err.Error())
// 		return
// 	}

// 	// Fetch user by ID from the service
// 	baseRes, user, err := u.userSvc.GetUserById(ctx, reqBody.ID)
// 	if err != nil {
// 		smerrors.InternalServer(ctx, err.Error())
// 		return
// 	}

// 	utils.ReturnJSONStruct(ctx, baseRes)
// }


// func (u *userHandler) DeleteUser(ctx *gin.Context) {
// 	// Validate the request body
// 	reqBody, err := validateDeleteUser(ctx)
// 	if err != nil {
// 		smerrors.Validation(ctx, err.Error())
// 		return
// 	}

// 	// Delete user in the service
// 	baseRes, err := u.userSvc.DeleteUser(ctx, reqBody.ID)
// 	if err != nil {
// 		smerrors.InternalServer(ctx, err.Error())
// 		return
// 	}

// 	utils.ReturnJSONStruct(ctx, baseRes)
// }