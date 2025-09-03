package models

import "time"

type Message struct {
	ID         uint   `gorm:"primary_key"`
	SenderID   uint   `gorm:"not null"`
	ReceiverID uint   `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
	IsRead     bool   `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
