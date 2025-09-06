package applications

import (
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/domain/models"
)

type UserUsecase struct {
	userRepo *repositories.UserRepository
}

func NewUserUsecase(userRepo *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) SearchUsers(query string) ([]*models.User, error) {
	return u.userRepo.SearchUsers(query)
}
