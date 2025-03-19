package models

import (
	"time"
)

type User struct {
	IdUser   uint   `gorm:"primary_key;autoIncrement;uniqueIndex;primary_key" json:"id_user"`
	Username string `gorm:"size:150;not null;uniqueIndex" json:"username"`
	Email    string `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Message []Message `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	IdUser    uint      `json:"id_user"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserCreate struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"johndoe@gmail.com"`
	Password string `json:"password" example:"passwordJohndoe"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
