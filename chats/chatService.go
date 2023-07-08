package chats

import (
	"Chat/chats/models"
	"context"
	"github.com/google/uuid"
)

type IChatService interface {
	IChatRepo
	NewChat(c context.Context, chat *models.Chat) (models.Chat, error)
}

type ChatService struct {
	IChatRepo
}

// NewChat
//
// Create a new chat
func (cs *ChatService) NewChat(c context.Context, chat *models.Chat) (models.Chat, error) {
	chat.Id = uuid.NewString()
	newChat, err := cs.CreateNew(c, *chat)
	return newChat, err
}

// NewChatService
//
// Returns the instance of ChatServices
func NewChatService(chatRepo IChatRepo) *ChatService {
	return &ChatService{
		IChatRepo: chatRepo,
	}
}
