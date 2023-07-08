package models

import (
	"Chat/base"
	"Chat/chats/models"
	userModel "Chat/users/models"
)

// Conversation
//
// represents a conversation between two users.
type Conversation struct {
	base.BaseModel // BaseModel contains common fields such as ID, CreatedAt, and UpdatedAt.

	// SenderId is the ID of the user who initiated the conversation.
	SenderId string `json:"senderId" gorm:"column:senderId"`

	// ReceiverId is the ID of the user who is receiving the conversation.
	ReceiverId string `json:"receiverId" gorm:"column:receiverId"`

	// Chats is a list of messages sent in this conversation.
	Chats []models.Chat `gorm:"foreignKey:ConversationId"`

	// Sender is the user who initiated the conversation.
	Sender userModel.User `json:"sender,omitempty" gorm:"foreignKey:SenderId"`

	// Receiver is the user who is receiving the convedrsation.
	Receiver userModel.User `json:"receiver,omitempty" gorm:"foreignKey:ReceiverId"`
}
