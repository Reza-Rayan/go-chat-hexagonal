package repositories

import (
	"github.com/Reza-Rayan/internal/domain/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	Db *gorm.DB
}

func (r *MessageRepository) SaveMessage(msg *models.Message) error {
	return r.Db.Create(msg).Error
}

func (r *MessageRepository) GetMessagesBetweenUsers(userID1, userID2 uint) ([]*models.Message, error) {
	var messages []*models.Message
	err := r.Db.Where(
		"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID1, userID2, userID2, userID1,
	).Order("created_at asc").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
