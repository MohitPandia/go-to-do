package user

import (
	"fmt"
	"go-to-do/smerrors"
	"go-to-do/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. Validate request
// 2. Send Request to service layer
// 3. Get response from service layer
// 4. Send Response back with transformers

// --> 3 things asked to do
// 1. make the delete soft
// 2. also give a option to update {use map string interface to update}
// 3. also in the get all user use pagination

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

/* -------------------------------------------------------------------------- */
/*                                 getAlGetAllUsers                           */
/* -------------------------------------------------------------------------- */

func (u *userHandler) GetAllUsers(ctx *gin.Context) {
	// Validate the request
	// Validate the request
	reqBody, err := validateGetAllUsers(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	// Extract pagination query params
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if page != "" {
		reqBody.Page, _ = strconv.Atoi(page)
	}
	if limit != "" {
		reqBody.Limit, _ = strconv.Atoi(limit)
	}

	// Call the service method
	baseRes, users, err := u.userSvc.GetAllUsers(ctx, reqBody)
	if err != nil {
		smerrors.InternalServer(ctx, err.Error())
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response if needed
	finalRes := GetAllUsertransformer(baseRes, users)
	utils.ReturnJSONStruct(ctx, finalRes)
}

/* -------------------------------------------------------------------------- */
/*                                GetUserByPID                                */
/* -------------------------------------------------------------------------- */

func (u *userHandler) GetUserByPID(ctx *gin.Context) {
	// Validate request body
	reqBody, err := validateGetUserByPID(ctx)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	fmt.Println("reqBody", reqBody)
	// Call the service method
	baseRes, user, err := u.userSvc.GetUserByPID(ctx, reqBody)
	if err != nil {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	fmt.Println("baseRes", baseRes)
	fmt.Println("user", user)
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response
	finalRes := TransformGetUserResponse(baseRes, user)
	utils.ReturnJSONStruct(ctx, finalRes)
}

/* -------------------------------------------------------------------------- */
/*                                 DeleteUser                                 */
/* -------------------------------------------------------------------------- */
func (u *userHandler) DeleteUser(ctx *gin.Context) {
	// Validate request body
	reqBody, err := validateDeleteUser(ctx)
	fmt.Println("reqBody in controller", reqBody)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	// Call the service method
	baseRes, users, err := u.userSvc.DeleteUser(ctx, reqBody)
	fmt.Println("baseRes in controller", baseRes)

	if err != nil {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Return success response
	finalRes := DeleteUserTransformer(baseRes, users)

	utils.ReturnJSONStruct(ctx, finalRes)
}

/* -------------------------------------------------------------------------- */
/*                                 UpdateUser                                 */
/* -------------------------------------------------------------------------- */
func (u *userHandler) UpdateUser(ctx *gin.Context) {
	// Validate request body

	reqBody, err := validateUpdateUser(ctx)
	fmt.Println("reqBody in controller", reqBody)
	if err != nil {
		smerrors.Validation(ctx, err.Error())
		return
	}

	fmt.Println("reqBody in controller", reqBody)

	// Call the service method
	baseRes, updatedUser, err := u.userSvc.UpdateUser(ctx, reqBody)
	if err != nil {
		smerrors.InternalServer(ctx, err.Error())
		return
	}
	if baseRes.StatusCode != http.StatusOK {
		smerrors.HandleServiceCodes(ctx, baseRes)
		return
	}

	// Transform the response
	finalRes := UpdateUserTransformer(baseRes, updatedUser)
	utils.ReturnJSONStruct(ctx, finalRes)
}
