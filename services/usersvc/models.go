package usersvc

// CreateUserRequest represents the request payload for creating a new user.
// type CreateUserRequest struct {
// 	Name     string `json:"name" validate:"required,min=3,max=100"`     // Name must be 3-100 characters long.
// 	Email    string `json:"email" validate:"required,email,max=100"`    // Email must be valid and max 100 characters.
// 	Password string `json:"password" validate:"required,min=8,max=100"` // Password must be 8-100 characters long.
// }

type CreateUserObject struct {
	Name      string `json:"name" validate:"required,min=3,max=100"`     // Name must be 3-100 characters long.
	Email     string `json:"email" validate:"required,email,max=100"`    // Email must be valid and max 100 characters.
	Password  string `json:"password" validate:"required,min=8,max=100"` // Password must be 8-100 characters long.
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

// GetAllUserObject serves as a placeholder for future enhancements,
// such as filtering or pagination parameters for fetching all users.
type GetAllUserObject struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// GetUserByPIDObject represents the request payload for fetching a user by PID.
type GetUserByPIDObject struct {
	PID string `json:"pid" validate:"required"` // PID must be exactly 40 characters long.
}

// DeleteUserObject represents the request payload for deleting a user.
type DeleteUserObject struct {
	PID string `json:"pid" validate:"required"` // User PID is required for deletion
}

type CreateUserRes struct {
	Name string `json:"name"`
}

// UpdateUserObject represents the request payload for updating a user's details.
type UpdateUserObject struct {
	PID      string  `json:"pid" validate:"required"`                   // User PID to identify the user being updated.
	Name     *string `json:"name" validate:"omitempty,min=3,max=100"`   // Optional: Name must be 3-100 characters long.
	Email    *string `json:"email" validate:"omitempty,email,max=100"`  // Optional: Email must be valid and max 100 characters.
	Password *string `json:"password" validate:"omitempty,min=8,max=100"` // Optional: Password must be 8-100 characters long.
}