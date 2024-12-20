package usersvc

// CreateUserRequest represents the request payload for creating a new user.
// type CreateUserRequest struct {
// 	Name     string `json:"name" validate:"required,min=3,max=100"`     // Name must be 3-100 characters long.
// 	Email    string `json:"email" validate:"required,email,max=100"`    // Email must be valid and max 100 characters.
// 	Password string `json:"password" validate:"required,min=8,max=100"` // Password must be 8-100 characters long.
// }


type CreateUserObject struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`     // Name must be 3-100 characters long.
	Email    string `json:"email" validate:"required,email,max=100"`    // Email must be valid and max 100 characters.
	Password string `json:"password" validate:"required,min=8,max=100"`  // Password must be 8-100 characters long.
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}


// // GetUserByPIDObject represents the request payload for fetching a user by PID.
// type GetUserByPIDObject struct {
//     UserID int `json:"user_id" validate:"required"`
// }

// // DeleteUserObject represents the request payload for deleting a user.
// type DeleteUserObject struct {
//     UserID int `json:"user_id" validate:"required"`
// }




type CreateUserRes struct {
	Name string `json:"name"`
}
