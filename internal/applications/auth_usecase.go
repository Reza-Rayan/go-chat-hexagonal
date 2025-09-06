package applications

import (
	"errors"
	"github.com/Reza-Rayan/internal/adapters"
	"github.com/Reza-Rayan/internal/domain/models"
	"github.com/Reza-Rayan/utils"
	"time"
)

type AuthUsecase struct {
	userRepo *adapters.UserRepository
}

func NewAuthUsecase(userRepo *adapters.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

// Signup -> POST method
func (auth *AuthUsecase) Signup(username, password, email, phone, avatar string) (*models.User, error) {
	// Check username exists or not
	if existing, _ := auth.userRepo.FindByUsername(username); existing != nil {
		return nil, errors.New("username already taken")
	}

	// Check email exists or not
	if existing, _ := auth.userRepo.FindUserByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}

	//	Create & HashPassword
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username:  username,
		Email:     email,
		Phone:     phone,
		Avatar:    avatar,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := auth.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""
	return createdUser, nil

}

// Login -> POST method
func (auth *AuthUsecase) Login(email, password string) (string, *models.User, error) {
	//	Find User
	user, err := auth.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	//	Check Password
	if !utils.CheckPassword(password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}

	// return user & token
	user.Password = ""

	return token, user, nil
}
