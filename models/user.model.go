package models

import (
	"go-todo-app/request"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int                          `json:"id" gorm:"primaryKey"`
	Name      string                       `json:"name"`
	Email     string                       `json:"email" gorm:"unique"`
	Password  string                       `json:"-" gorm:"column:password"`
	Address   string                       `json:"address"`
	Phone     string                       `json:"phone"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt time.Time                    `json:"update_at"`
	DeletedAt gorm.DeletedAt               `json:"-" gorm:"index,column:deleted_at"`
	Todos     []request.TodoUpdateeRequest `json:"todos"`
}
