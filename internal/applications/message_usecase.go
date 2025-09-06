package applications

import (
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/domain/models"
	"time"
)

type MessageUsecase struct {
	messageRepo *repositories.MessageRepository
}

func NewMessageUsecase(repo *repositories.MessageRepository) *MessageUsecase {
	return &MessageUsecase{
		messageRepo: repo,
	}
}

func (uc *MessageUsecase) SendMessage(SenderID, ReceiverID uint, content string) (*models.Message, error) {
	msg := &models.Message{
		SenderID:   SenderID,
		ReceiverID: ReceiverID,
		Content:    content,
		IsRead:     false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := uc.messageRepo.SaveMessage(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (uc *MessageUsecase) GetMessages(userID1, userID2 uint) ([]*models.Message, error) {
	return uc.messageRepo.GetMessagesBetweenUsers(userID1, userID2)
}
