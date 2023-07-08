package fileUpload

import (
	"Chat/base"
	"Chat/fileUpload/models"

	"gorm.io/gorm"
)

type IAttachmentRepo interface {
	base.IBaseRepo[models.Attachment]
}

type AttachmentRepo struct {
	base.BaseRepo[models.Attachment]
}

func NewFileRepo(db *gorm.DB) *AttachmentRepo {
	return &AttachmentRepo{
		BaseRepo: base.BaseRepo[models.Attachment]{
			Db: db,
		},
	}
}
