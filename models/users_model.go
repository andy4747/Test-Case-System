package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	Superuser bool   `json:"superuser"`
}

func (Users) TableName() string {
	return "users"
}
