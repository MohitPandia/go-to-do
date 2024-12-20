package entities

import (
	"time"
)

type Users struct {
	ID        int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID       string    `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	Name      string    `gorm:"column:name;not null;type:varchar(100)"`
	Email     string    `gorm:"column:email;unique;not null;type:varchar(100)"`
	Password  string    `gorm:"column:password;not null;type:varchar(100)"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
