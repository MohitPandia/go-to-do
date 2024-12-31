package entities

import (
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	ID          int            `gorm:"column:todo_id;primaryKey;autoIncrement"`
	PID         string         `gorm:"column:todos_pid;unique;not null;type:varchar(40)"`
	UserPID     string         `gorm:"column:user_pid;not null;type:varchar(40)"`
	CategoryPID string         `gorm:"column:category_pid;not null;type:varchar(40)"`
	Title       string         `gorm:"column:title;not null;type:varchar(200)"`
	Description string         `gorm:"column:description;type:text"`
	DueDate     time.Time      `gorm:"column:due_date"`
	Completed   bool           `gorm:"column:completed;not null;default:false"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"` // Adds support for soft deletes
}
