package fileUpload

import (
	"Chat/base"
	"Chat/fileUpload/models"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type IAttachmentService interface {
	IAttachmentRepo
	SaveFile(file *multipart.FileHeader, id string, ctx context.Context) (models.Attachment, error)
	SaveMultipleFile() ([]models.Attachment, error)
}

type AttachmentService struct {
	IAttachmentRepo
}

// SaveFile
//
// Saves the given file in folder required
// new goroutine is instantiated for it
func (f *AttachmentService) SaveFile(file *multipart.FileHeader, id string, ctx context.Context) (models.Attachment, error) {
	attachment := models.Attachment{
		Name:   file.Filename,
		Path:   "data/attachment/", //Todo generate based on content
		UserId: id,
		BaseModel: base.BaseModel{
			Id: uuid.NewString(),
		},
	}
	_, err2 := f.CreateNew(ctx, attachment)
	if err2 != nil {
		return models.Attachment{}, err2
	}
	src, err := file.Open()
	if err != nil {
		return models.Attachment{}, err
	}
	go SaveFileInBackground(src, file)

	return models.Attachment{}, nil //Todo add to database
}

func (f *AttachmentService) SaveMultipleFile() ([]models.Attachment, error) {
	return []models.Attachment{}, nil //Todo
}

func NewFileService(fileRepo IAttachmentRepo) *AttachmentService {
	return &AttachmentService{
		IAttachmentRepo: fileRepo,
	}
}

/*
#########################################################################################
									Helpers
#########################################################################################
*/

// SaveFileInBackground
//
// Function to be run asynchronously
func SaveFileInBackground(file multipart.File, header *multipart.FileHeader) {
	defer file.Close()

	dst, err := os.Create("data/attachment/" + header.Filename) //Todo need to change the data  folder to env
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dst.Close() //Todo
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println(err.Error())
		return
	}
}
