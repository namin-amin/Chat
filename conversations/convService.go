package conversations

import (
	"Chat/conversations/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type IConvService interface {
	IConvRepo
	CreateNewConv(senderId string, receiverId string) (models.Conversation, error)
	GetConv(senderReceiverId string) ([]models.Conversation, error)
	GetConvWithPagination(ctx context.Context, itemCount int, numberOfPages int, lastRowId string) ([]models.Conversation, error)
}

type ConvService struct {
	IConvRepo
}

func (s *ConvService) CreateNewConv(senderId string, receiverId string) (models.Conversation, error) {
	var conv models.Conversation
	s.GetDb().Find(&conv, "senderId=? AND receiverId=?", senderId, receiverId)
	if conv.Id != "" {
		return conv, nil
	}
	conv.Id = uuid.NewString()
	conv.ReceiverId = receiverId
	conv.SenderId = senderId
	returnVal := s.GetDb().Create(&conv)

	if returnVal.Error != nil {
		return conv, returnVal.Error
	}
	s.GetDb().Preload(clause.Associations).Find(&conv, "id=?", conv.Id)
	return conv, nil
}

func (s *ConvService) GetConv(senderReceiverId string) ([]models.Conversation, error) {
	var conv []models.Conversation
	returnVal := s.GetDb().Preload(clause.Associations).Find(&conv, "senderId=? OR receiverId=?", senderReceiverId, senderReceiverId)

	if returnVal.Error != nil {
		return conv, returnVal.Error
	}

	return conv, nil
}

func (s *ConvService) GetConvWithPagination(ctx context.Context, itemCount int, numberOfPages int, lastRowId string) ([]models.Conversation, error) {
	return s.GetWithPagination(ctx, itemCount, numberOfPages, lastRowId)
}

func NewConvService(convRepo IConvRepo) *ConvService {
	return &ConvService{
		IConvRepo: convRepo,
	}
}
