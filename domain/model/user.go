package model

import "time"

type User struct {
	ID uint `gorm:"primary_key" json:"id"`
	UUID string `json:"uuid"`
	Email string `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}


func (u *User) TableName() string {
	return "users"
}