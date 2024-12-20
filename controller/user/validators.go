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
