package applications

import (
	"errors"
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/domain/models"
)

type FriendUsecase struct {
	userRepo *repositories.UserRepository
}

func NewFriendUsecase(userRepo *repositories.UserRepository) *FriendUsecase {
	return &FriendUsecase{userRepo: userRepo}
}

// GetFriends -> GET method
func (f *FriendUsecase) GetFriends(userID uint) ([]*models.User, error) {
	user, err := f.userRepo.FindByIDWithFriends(userID)
	if err != nil {
		return nil, err
	}
	return user.Friends, nil
}

// AddFriend -> POST method
func (f *FriendUsecase) AddFriend(userID, friendID uint) error {
	if userID == friendID {
		return errors.New("cannot add yourself as friend")
	}
	return f.userRepo.AddFriend(userID, friendID)
}
