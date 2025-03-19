package models

import "time"

type Message struct {
	IdMessage uint `gorm:"primary_key;autoIncrement" json:"id_message"`
	UserId    uint `gorm:"not null;" json:"user_id"`
	User      User `gorm:"foreignKey:UserId"`

	ConversationId uint         `gorm:"not null;" json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationId"`
	Content        string       `gorm:"type:text" json:"content"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type MessageResponse struct {
	IdMessage      uint      `json:"id_message"`
	UserId         uint      `json:"user_id"`
	ConversationId uint      `json:"conversation_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"sent_at"`
}

type MessageCreate struct {
	UserId  uint   `json:"user_id" example:"1"`
	Content string `json:"content" example:"Hai Apa kabar?"`
}

type ErrorMsgResponse struct {
	Error string `json:"error"`
}
