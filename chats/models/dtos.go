package models

import (
	msgModel "Chat/sse"
)

type NewChatDto struct {
	msgModel.Message
	ConversationId string `json:"conversationId"`
}
