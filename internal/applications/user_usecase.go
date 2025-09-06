package applications

import (
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/domain/models"
	"time"
)

type UserUsecase struct {
	userRepo *repositories.UserRepository
}

func NewUserUsecase(userRepo *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) SearchUsers(query string, limit, offset int) ([]*models.User, error) {
	return u.userRepo.SearchUsers(query, limit, offset)
}

// UpdateUser -> PUT method
func (u *UserUsecase) UpdateUser(userID uint, username, email, phone, avatarPath string) (*models.User, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
	if phone != "" {
		user.Phone = phone
	}
	if avatarPath != "" {
		user.Avatar = avatarPath
	}

	user.UpdatedAt = time.Now()

	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// FindUserByID -> Get Method
func (u *UserUsecase) FindUserByID(userID uint) (*models.User, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
