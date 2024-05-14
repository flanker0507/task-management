package models

import (
	"go-todo-app/request"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID         uint                 `json:"id" gorm:"primary_key"`
	Name       string               `json:"name" gorm:"not null"`
	Note       *string              `json:"note" gorm:""`
	IsComplete bool                 `json:"is_complete" gorm:"boolean,default:false"`
	UserId     uint                 `json:"user_id"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	DeletedAt  gorm.DeletedAt       `json:"-"`
	User       request.UserResponse `json:"user"`
}
