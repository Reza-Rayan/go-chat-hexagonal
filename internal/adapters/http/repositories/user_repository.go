package repositories

import (
	"github.com/Reza-Rayan/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.Db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByIDWithFriends(id uint) (*models.User, error) {
	var user models.User
	if err := r.Db.Preload("Friends").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AddFriend(userID, friendID uint) error {
	var user, friend models.User
	if err := r.Db.First(&user, userID).Error; err != nil {
		return err
	}
	if err := r.Db.First(&friend, friendID).Error; err != nil {
		return err
	}
	return r.Db.Model(&user).Association("Friends").Append(&friend)
}

func (r *UserRepository) SearchUsers(query string) ([]*models.User, error) {
	var users []*models.User
	if err := r.Db.Where("username LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	for _, u := range users {
		u.Password = ""
	}
	return users, nil
}
