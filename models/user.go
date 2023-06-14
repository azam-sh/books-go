package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	FullName  string         `json:"full_name" gorm:"column:full_name"`
	RoleID    int64          `json:"role_id" gorm:"column:role_id"`
	Login     string         `json:"login" gorm:"column:login;unique"`
	Password  string         `json:"password" gorm:"column:password"`
	Phone     string         `json:"phone" gorm:"column:phone"`
	Active    bool           `json:"active" gorm:"column:active"`
	LoginAt   *time.Time     `json:"login_at" gorm:"column:login_at"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Role      Role           `gorm:"foreignKey:RoleID"`
}

type ResponseUser struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullName"`
	RoleID   int64  `json:"roleId"`
	Phone    string `json:"phone"`
	Active   bool   `json:"active"`
}

type RequestUser struct {
	FullName string `json:"fullName"`
	RoleID   int64  `json:"roleId"`
	Phone    string `json:"phone"`
	Active   bool   `json:"active"`
}

func (u User) TableName() string {
	return "users"
}
