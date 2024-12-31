package todosvc

import "time"

type CreateTodoObject struct {
	UserPID     string    `json:"user_pid" validate:"required"`
	CategoryPID string    `json:"category_pid" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"descripton" validate:"required"`
	DueDate     time.Time `json:"dueDate" validate:"required"`
	Completed   bool      `json:"completed" validate:"required"`
	CreatedAt   int       `json:"created_at"`
	UpdatedAt   int       `json:"updated_at"`
}


type GetTodoByPIDObject struct {
	PID string `json:"pid" validate:"required"` // PID must be exactly 40 characters long.
}