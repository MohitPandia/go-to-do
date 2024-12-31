package user

import (
	"errors"
	"fmt"
	"go-to-do/services/usersvc"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/* -------------------------------------------------------------------------- */
/*                    Validate Create KnowledgeBase Request                   */
/* -------------------------------------------------------------------------- */
func validateCreateUser(c *gin.Context) (usersvc.CreateUserObject, error) {
	var reqBody usersvc.CreateUserObject
	fmt.Println("in validatior", reqBody)
	var err error

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	// Trim the name and validate it
	reqBody.Name = strings.TrimSpace(reqBody.Name)
	err = validateName(reqBody.Name)
	if err != nil {
		return reqBody, err
	}

	validate := validator.New()
	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	fmt.Println("checking that the space is trimmerd", reqBody.Name)

	return reqBody, err
}

// validateName validates the user name for specific criteria.
func validateName(name string) error {
	// Trim leading and trailing whitespace
	trimmedName := strings.TrimSpace(name)

	// Check if the trimmed name is empty
	if len(trimmedName) == 0 {
		return errors.New("name cannot be empty or just whitespace")
	}

	// Validate length constraints
	if len(trimmedName) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	if len(trimmedName) > 50 {
		return errors.New("name must be no more than 50 characters long")
	}

	return nil
}

func validateGetAllUsers(ctx *gin.Context) (usersvc.GetAllUserObject, error) {
	var reqBody usersvc.GetAllUserObject
	var err error

	fmt.Println("in validatior", reqBody)

	// Bind and validate JSON payload
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	validate := validator.New()
	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	fmt.Println(reqBody, "reqBody")
	return reqBody, nil
}

func validateGetUserByPID(ctx *gin.Context) (usersvc.GetUserByPIDObject, error) {
	var reqBody usersvc.GetUserByPIDObject
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	// Validate using the validator package
	validate := validator.New()
	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	return reqBody, nil
}

func validateDeleteUser(c *gin.Context) (usersvc.DeleteUserObject, error) {
	var reqBody usersvc.DeleteUserObject
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, err
	}

	// Validate that PID is not empty
	if reqBody.PID == "" {
		return reqBody, errors.New("user PID is required")
	}

	return reqBody, nil
}

func validateUpdateUser(ctx *gin.Context) (map[string]interface{}, error) {
	var reqBody map[string]interface{}
	var err error

	fmt.Println("coming in this validator")

	// Bind JSON payload into a map
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		return nil, err
	}

	// Initialize a validator instance
	validate := validator.New()

	// Required field validation for "PID"
	if pid, ok := reqBody["pid"].(string); !ok || strings.TrimSpace(pid) == "" {
		return nil, errors.New("user PID is required and must be a non-empty string")
	}

	// Optional field validation
	if name, ok := reqBody["name"].(string); ok {
		if err := validateName(strings.TrimSpace(name)); err != nil {
			return nil, err
		}
	}

	// Additional validations as needed
	if email, ok := reqBody["email"].(string); ok {
		if err := validate.Var(email, "email"); err != nil {
			return nil, errors.New("invalid email format")
		}
	}

	return reqBody, nil
}
