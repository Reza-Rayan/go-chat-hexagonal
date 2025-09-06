package handlers

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authUC *applications.AuthUsecase
}

func NewAuthHandler(authUC *applications.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

// Signup -> POST method
func (h *AuthHandler) Signup(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authUC.Signup(req.Username, req.Password, req.Email, req.Phone, req.Avatar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully", "user": user})
}

// Login -> POST method
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.authUC.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"user":    user,
	})
}
