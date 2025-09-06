package models

import "time"

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"uniqueIndex;not null"`
	Password  string  `gorm:"not null"`
	Email     string  `gorm:"uniqueIndex;not null"`
	Phone     string  `gorm:"uniqueIndex"`
	Friends   []*User `gorm:"many2many:user_friends;" json:"friends,omitempty"`
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
