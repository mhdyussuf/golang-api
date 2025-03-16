package models

import "time"

type Conversation struct {
	IdConversation uint      `gorm:"primary_key;autoIncrement" json:"id_message"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Messages []Message `gorm:"foreignKey:ConversationId;constraint:OnDelete:CASCADE"`
}
