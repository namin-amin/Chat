package chats

import (
	"Chat/base"
	"Chat/chats/models"

	"gorm.io/gorm"
)

type IChatRepo interface {
	base.IBaseRepo[models.Chat]
}

type ChatRepo struct {
	base.BaseRepo[models.Chat]
}

// NewChatRepo
//
// Returns the instance of ChatRepository
func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		BaseRepo: base.BaseRepo[models.Chat]{
			Db: db,
		},
	}
}
