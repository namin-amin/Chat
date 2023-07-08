package models

import (
	"Chat/base"
	messageModel "Chat/sse"
)

type ChatType int

const (
	Text  ChatType = 1 //message with text data
	Media ChatType = 2 //message with media type of data like doc,pdf,png etc.
)

// Chat
//
// structure represents individual Chat
type Chat struct {
	base.BaseModel
	messageModel.Message        //message received or to be sent
	ConversationId       string `json:"conversationId" gorm:"column:conversationId"` // id of the conversation this chat belongs to
}
