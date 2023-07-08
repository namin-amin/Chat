package conversations

import (
	"Chat/base"
	"Chat/conversations/models"

	"gorm.io/gorm"
)

type IConvRepo interface {
	base.IBaseRepo[models.Conversation]
}

type ConvRepo struct {
	base.BaseRepo[models.Conversation]
}

func NewConvRepo(db *gorm.DB) *ConvRepo {
	return &ConvRepo{
		BaseRepo: base.BaseRepo[models.Conversation]{
			Db: db,
		},
	}
}
